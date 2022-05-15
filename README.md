This project shows how it can be difficult to encode multi-line strings for substitution in
functions project.yml.  It can be run against app platform with the [provided spec file](spec.yaml)
or the example can be run against `nim project:get-metadata` with the provided [json encoded var file](multiline.json).

```
nim project:get-metadata --env multiline.json .
{
  "environment": {
    "PARENTHESIS_MULTILINE": "line1\nline2"
  },
  "packages": [
    {
      "name": "vars",
      "actions": [
        {
          "name": "vars",
          "file": "packages/vars/vars",
          "displayFile": "/home/cbaker/DevRoot/funcvars/vars/vars",
          "build": "build.sh",
          "package": "vars",
          "web": true,
          "runtime": "go:default"
        }
      ],
      "shared": false,
      "environment": {
        "QUOTED_CURLY_BRACKET": "line1 line2",
        "INTERNALLY_INDENTED_CURLY_BRACKET_MULTILINE": "line1\nline2",
        "QUOTED_DOUBLE_ESCAPED_CURLY_BRACKET": "line1\nline2"
      }
    }
  ],
  "strays": [
    ".git",
    "README.md",
    "multiline.json"
  ],
  "filePath": ".",
  "includer": {
    "isWebIncluded": true,
    "isExcludingOnly": true,
    "includedPackages": {},
    "excludedPackages": {},
    "includedActions": {},
    "excludedActions": {}
  },
  "reader": {
    "basepath": "."
  },
  "feedback": {},
  "deployerAnnotation": {
    "user": "cbaker",
    "projectPath": "/home/cbaker/DevRoot/funcvars"
  },
  "web": []
}
```