![lima](https://github.com/faizisyellow/lima/blob/main/demo/lima-banner.png?raw=true)

# What is it ?
Lima is a movie list management command line interface application.  
With lima you can manage your movies list easyly and as simple as possible.

# Features
- Add a new movie to the list
- Read a list of movies
- Update a movie
- Mark a movie as watched
- Update current watching-movie's duration
- Delete a movie


# Installation
> Currently only manual installation is available, so lima requires [Go](https://go.dev/doc/install) to be installed

Clone lima repository

Get depedencies that lima needed

```go mod tidy ```

Build lima project :

``` make lima-build ```

Run:

```./bin/lima```

# Configuration
> Lima stores (json) list of movies through env variable and with prefix LIMA_

- Bash
    - Create a json file to store the list:

    ```touch $HOME/.lima_store.json```

    - Open your bashrc config, then export ENV variable:
    
    ``` export LIMA_store=$HOME/.lima_store.json``` 


# Usage Examples
An example of lists movies  

```
lima list
```  

![lima](https://github.com/faizisyellow/lima/blob/main/demo/lima-ls-demo.png?raw=true)

An example of add a new movie

```
lima add [title] --year --category --status
```  

![lima](https://github.com/faizisyellow/lima/blob/main/demo/lima-add-demo.png?raw=true)

An example of edit a movie


```
lima edit [id] --title --year
```  

![lima](https://github.com/faizisyellow/lima/blob/main/demo/lima-update-demo.png?raw=true)

An example of remove a movie


```
lima remove [id]
```  

![lima](https://github.com/faizisyellow/lima/blob/main/demo/lima-rm-demo.png?raw=true)

Use --help to see all other commands

``` lima --help ```