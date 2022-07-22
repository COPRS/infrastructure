# Add your content packs to graylog

## Purpose

Graylog has a mecanism that allows administrators to export a part of their graylog configuration's.
That's include Inputs, Extractors, Pipelines, Gork Patterns etc..
https://www.graylog.org/features/content-packs

In a cloud context, that is helpfull as platforms can be redeployed many times.
This document explains how you can easily import your graylog content pack.

## Update kustomization.yaml
The first step, consist in create a folder `content-packs` in the `apps/graylog` folder.
Inside this folder you can copy past your content packs in json format.

Once done, you have to update the `kustomization.yaml` file of Graylog.
The purpose here is to add as `ConfigMap` resources the content packs located in your folder.

```
configMapGenerator:
  - name: graylog-contentpacks
    files:
      - ./content-packs/contentpacks.json
```

## Update values.yaml

Now that the content packs are easily accessible by graylog, all we have to do is to
update the `values.yaml` to mount those files inside the graylog server container and then to update the `graylog.conf` to install them at the graylog deployment.

Under the graylog object in values.yaml adapts the munder to
```
  extraVolumeMounts:
    - name: graylog-contentpacks
      mountPath: /usr/share/graylog/data/content-packs
  extraVolumes:
    - name: graylog-contentpacks
      configMap: 
        name: graylog-contentpacks
  config: |
    content_packs_loader_enabled = true
    content_packs_dir = /usr/share/graylog/data/content-packs/
    content_packs_auto_install = contentpacks.json,grok-patterns.json
```
The values add to `config` tag are added at the of the `graylog.conf` file inside the graylog container.