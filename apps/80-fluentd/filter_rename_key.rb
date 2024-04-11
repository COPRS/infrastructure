# Copyright 2023 CS Group
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

# source: https://github.com/shunwen/fluent-plugin-rename-key/blob/master/lib/fluent/plugin/filter_rename_key.rb

require 'fluent/plugin/filter'
require '/etc/fluent/custom_plugins/rename_key_util'

class Fluent::Plugin::RenameKeyFilter < Fluent::Plugin::Filter
  Fluent::Plugin.register_filter 'rename_key', self

  include Fluent::Plugin::RenameKeyUtil

  desc 'Deep rename/replace operation.'
  config_param :deep_rename, :bool, default: true

  def configure conf
    super

    create_rename_rules(conf)
    create_replace_rules(conf)

    if @rename_rules.empty? && @replace_rules.empty?
      raise Fluent::ConfigError, 'No rename nor replace rule given'
    end
  end

  def filter _tag, _time, record
    replace_key(rename_key(record))
  end
end