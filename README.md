A Boot dot Dev project to build a Pokedexcli tool in Go.

The purpose of the project is to get practice using http.GET to get information about Pokemon from the Pokemon website. This includes parsing json files and structuring Go projects in github.

The project has a simple structure:
The main package:
    main.go has an infinite loop that scans for user keystrokes
    repl.go has a map of the commands the user can call and the function definitions of exit and help commands

Internal packages
    pokeapi:
    maps.go has map and mapb functions that display a page of 20 pokemon location-areas. map displays the next page of 20 locations, mapb will display the previous page of 20 locations.
    explore.go has the explore function that lists the pokemon in a location entered by the user
    catch.go has the catch function where the user can try to catch a pokemon in an area they have explored.
    inspect.go has the inspect function that displays attributes of a pokemon entered by the user from the pokemon that the user has already caught.
    pokedex.go has the pokedex function that displays the names of all the pokemon that the user has caught.

    pokecache:
    pokecache.go has all the caching logic and functions.
