# Ebitenengine playground
This repository contains results of my first trial with Ebitengine :)

The code consists of both server and client code of a very simple concept of a top-down
pixel-art styled multiplayer game. In current state it just let you log-in to a locally running
server and just move around a very first and unfinished world scene.

Client details:

* Built on top of [Ebitengine](https://ebitengine.org/),
* Uses ECS (Entity-Component-System) provided by [Donburi](https://github.com/yottahmd/donburi-ecs)
* GUI done with [Furex] (https://github.com/yottahmd/furex-ui)
* Tiled background edited with [Tiled](https://www.mapeditor.org/)
* Path finding done with [go-astar](https://github.com/nickdavies/go-astar)

Game assets, including character with animations, trees, grass, rivers, and 
basically everything you can find inside comes from https://game-endeavor.itch.io/mystic-woods
Because of the licensing model I cannot include png files directly in this repo - you need to get them yourself.
Then by looking at .tsx files you can try to think where to put them ;) [or ask me....]

Some server insides:
* Provides HTTP controllers for handling user sessions,
* Does some imperfect UDP to synchronize remote users positions (meh...)
* Uses PostgreSQL to store user credentials,
* [GORM](https://gorm.io/index.html) for ORM,
* Kinda follows DDD...

You can use docker-compose.yml file located under docker directory to spin up the DB locally.

There are two users available in DB migration scripts:
* user: test, pass: test123
* user: test2, pass: test123

(yeah, security...)