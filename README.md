gopher repo 

-Submodules 
 - when tagging modules vesion the prefi should contain the path from the root
 e.g if module is at github.com/OdinPogi/gopher/modules/submodules/b and root is 
 github.com/OdinPogi/gopher

 tagging should be as follows : 

git tag modules/submodules/b/v1.0.0
$ git push -q origin modules/submodules/b/v1.0.0

+Cleam module cache:
go clean -cache -modcache -i -r
