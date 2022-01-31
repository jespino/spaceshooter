## Space Shooter

This is a port to go of the pygame Space Shooter (https://github.com/tasdikrahman/spaceShooter) game using the ebiten library.

## Index

- [Demo](https://github.com/jespino/spaceshooter#demo)
  - [Screenshots](https://github.com/jespino/spaceshooter#screenshots)
- [Game Features](https://github.com/jespino/spaceshooter#game-features)
  - [Controls](https://github.com/jespino/spaceshooter#controls)
- [Installation](https://github.com/jespino/spaceshooter#installation)
  - [Option 1: Download the zipped executable file](https://github.com/jespino/spaceshooter#option-1-download-the-zipped-executable-file)
  - [Option 2: Build from source](https://github.com/jespino/spaceshooter#option-2-build-from-source)
- [To-do](https://github.com/jespino/spaceshooter#to-do)
- [Issues](https://github.com/jespino/spaceshooter#issues)
- [Credits](https://github.com/jespino/spaceshooter#credits)
- [Similar](https://github.com/jespino/spaceshooter#similar)
- [License](https://github.com/jespino/spaceshooter#license)

## Demo

[[Back to top]](https://github.com/jespino/spaceshooter#index)

You can directly play this using web assambly compiled version here: https://jespino.github.io/spaceshooter/

## Screenshots

[[Back to top]](https://github.com/jespino/spaceshooter#index)

| ![Screen 1](http://i.imgur.com/3MzfmbT.jpg) | ![Screen 2](http://i.imgur.com/4OgIByR.png) |
|---------------------------------------------|---------------------------------------------|
| ![Screen 3](http://i.imgur.com/PFQJjE8.png) | ![Screen 4](http://i.imgur.com/lV4aIur.png) |

## Game Features

[[Back to top]](https://github.com/jespino/spaceshooter#index)

- Health bar for the space ship
- Score board to show how you are faring so far
- Power ups like
  - shield: increases the space ships life
  - bolt: increases the shooting capability of the ship by firing 2 to 3 bullets instead of one at time.
- Custom sounds and sprite animation for things like
  - meteorite explosion
  - bullet shoots
  - player explosion
- 3 lives per game
- Fun to play :)

## Controls

[[Back to top]](https://github.com/jespino/spaceshooter#index)

|              | Button              |
|--------------|---------------------|
| Move Left    | <kbd>left</kbd>     |
| Move right   | <kbd>right</kbd>    |
| Fire bullets | <kbd>spacebar</kbd> |
| Quit game    | <kbd>q</kbd>        |

## Installation

[[Back to top]](https://github.com/jespino/spaceshooter#index)

#### Option 1: Download the executable file

- :arrow_down: [Download the latest file for your operating system](https://github.com/jespino/spaceshooter/releases/latest)

If your download was saved on the `~/Downloads` folder

```bash
$ cd ~/Downloads
~/Downloads $ chmod +x spaceshooter
~/Downloads $ ./spaceshooter
```

#### Option 2: Build from source

```sh
$ git clone https://github.com/jespino/spaceshooter.git
$ cd spaceshooter/
$ go build -o ./ ./...
$ ./spaceshooter
```

### To-do

[[Back to top]](https://github.com/jespino/spaceshooter#index)

- [ ] Add feature to pause to the game.
- [ ] add feature to replay the game after all players die

## Issues

[[Back to top]](https://github.com/jespino/spaceshooter#index)

You can report the bugs at the [issue tracker](https://github.com/jespino/spaceshooter/issues)

## Credits

The original game is a fork of the video instructions given by KidsCanCode. I have made several additional enhancements to it. Do check out their [Channel](https://www.youtube.com/channel/UCNaPQ5uLX5iIEHUCLmfAgKg)!

## License

[[Back to top]](https://github.com/jespino/spaceshooter#index)

The original version was build by [Tasdik Rahman](http://tasdikrahman.me)[(@tasdikrahman)](https://twitter.com/tasdikrahman) under [MIT License](http://tasdikrahman.mit-license.org)

This port to go is build by Jesús Espino under [MIT License](http://mit-license.org)

- The images used in the game are taken from [http://opengameart.org/](http://opengameart.org/), more particulary from the [Space shooter content pack](http://opengameart.org/content/space-shooter-redux) from [@kenney](http://opengameart.org/users/kenney).

License for them is in `Public Domain`

- The game sounds were again taken from [http://opengameart.org/](http://opengameart.org/). The game music, [Frozen Jam](http://opengameart.org/content/frozen-jam-seamless-loop) by [tgfcoder](https://twitter.com/tgfcoder) licensed under [CC-BY-3](http://creativecommons.org/licenses/by/3.0/)
