# go-class-dep-mgmt-mod-b

Module used as part of an example to show how dependency management works - Module B

This module exposes 3 APIs in 3 different packages:

- an API that calls the API exposed by module C that calls an API exposed by module D
- an API that calls the API exposed by module C that remains internal to module C
- an API that does not call any other external API
