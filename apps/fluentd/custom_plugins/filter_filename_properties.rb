#
# Copyright 2019- CS Group
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

require "fluent/plugin/filter"
require "date"

module Fluent
    module Plugin
        class FilenamePropertiesFilter < Fluent::Plugin::Filter
            Fluent::Plugin.register_filter("filename_properties", self)

            config_param :key, :string, :default => nil

            def filter(tag, time, record)
                log=nil
                if !@key.nil? && record.has_key?(@key)
                    log = record[@key]
                else 
                    log = record
                end
                unless log.nil? || !log.has_key?("task")
                    task = log["task"]
                    modify_task_on(task, "input")
                    modify_task_on(task, "output")
                end
                record
            end

            private

            def modify_task_on(task, ioput)
              if task.has_key?(ioput) 
                if task[ioput].has_key?("filename_strings")
                  task[ioput]["filenames"] = transform_list(task[ioput]["filename_strings"])
                  task[ioput].delete("filename_strings")
                end
                if task[ioput].has_key?("segment_strings")
                  task[ioput]["segments"] = transform_list(task[ioput]["segment_strings"])
                  task[ioput].delete("segment_strings")
                end
              end
            end

            def transform_list(filenames)
              new_filenames = []
              filenames.each do |filename|
                new_filenames.push(transform(filename))
              end
              new_filenames
            end

            def transform(filename)
                pattern = %r{
                    ^(?<mission_identifier>S1[A-C])
                    _(?<mode_beam_identifier>[[:alnum:]]{2})
                    _(?<product_type>[[:alnum:]]{3})
                    ((?<resolution_class>[FHM])|_)
                    _(?<processing_level>[0-2])
                    ((?<product_class>[SACN])|_)
                    ((?<polarisation>[SDHV][HV])|__)
                    _(?<start>\d{8}T\d{6})
                    _(?<stop>\d{8}T\d{6})
                    _(?<absolute_orbite_number>\d{6})
                    _(_{6}|(?<mission_data_take_id>[[:xdigit:]]{6}))
                    _(?<product_unique_id>[[:xdigit:]]{0,4})
                    \.(?<product_file_extension>SAFE)$
                }x
                element = Hash["filename" => filename]
                if filename.end_with?(".SAFE")
                    # We can't use named_captures because fluentd uses ruby < 2.4
                    # element.update(filename.match(pattern).named_captures)
                    match = filename.match(pattern)
                    unless match.nil?
                        hmatch = match.names.zip(match.captures).to_h
                        hmatch.delete_if{ |k,v| v.nil? || v.empty? }
                        hmatch["start"] = DateTime.parse(hmatch["start"]).strftime("%FT%T.%6NZ")
                        hmatch["stop"] = DateTime.parse(hmatch["stop"]).strftime("%FT%T.%6NZ")
                        hmatch["processing_level"] = hmatch["processing_level"].to_i
                        hmatch["absolute_orbite_number"] = hmatch["absolute_orbite_number"].to_i
                        element.update(hmatch)
                    end
                end
                element
            end
        end
    end
end
