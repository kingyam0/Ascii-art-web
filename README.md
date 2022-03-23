    _____________________________________   _______       _____     ___       __    ______  
    ___    |_  ___/_  ____/___  _/___  _/   ___    |________  /_    __ |     / /_______  /_ 
    __  /| |____ \_  /     __  /  __  /     __  /| |_  ___/  __/    __ | /| / /_  _ \_  __ \
    _  ___ |___/ // /___  __/ /  __/ /      _  ___ |  /   / /_      __ |/ |/ / /  __/  /_/ /
    /_/  |_/____/ \____/  /___/  /___/      /_/  |_/_/    \__/      ____/|__/  \___//_.___/ 

Team (Authors):
- Gulim
- Kingsley
- Nik


# [Ascii-art-web ](https://github.com/01-edu/public/tree/master/subjects/ascii-art-web)
consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) for the [ASCII art project](https://github.com/01-edu/public/tree/master/subjects/ascii-art)

# Usage
Endpoints:
- `/` Home page     GET
- `/ascii-art`      POST

![webserver endpointsplan](https://i.imgur.com/MpNEMk4.jpg)

# Status Codes implemented
- OK (200), if everything went without errors
- Not Found, if nothing is found, for example templates or banners.
- Bad Request, for incorrect requests.
- Internal Server Error, for unhandled errors.

![status code plan](https://i.imgur.com/IhpQVfD.jpg)

# To Do:
- ~~When clicking anywhere on the text area.. make the cursor start from the begining ~~ :white_check_mark: 
- handle error if text contains unwanted characters :white_check_mark: 
- clean up and refactor :white_check_mark: 

## maybe to do:
implement the project to run using an API?