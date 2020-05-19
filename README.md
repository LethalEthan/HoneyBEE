# HoneyGO

HoneyGO is a Minecraft server written in GO, with help from [wiki.vg](https://wiki.vg) (Huge thanks to them and I suggest helping them out with a donation for the awesome resources that they publish for free <3).

This software is still being developed so expect bugs and issues now and then, also remember that feedback and disclosure of issues is welcomed but please don't send any hateful or spam issues :). HoneyGO as of now is considered pre-alpha it is very early in development and has only been released publicly on github since 25/4/2020

**Disclaimer: All code is completely original and is not based off of Mojang's code, any PR's must not have any of Mojang's copyrighted code as it will not be accepted nor be encouraged because it would be illegal, please see the Terms and Conditions of Mojang for familiarisation. I am not in any way affiliated with Mojang, I'm just a kid who wants a fast non-laggy mc server :) and because of that this is a learning curve for me and I hope I inspire other people to get into programming and reverse engineering as I am having a lot of fun making this. \
\
All code is commented and is aimed at beginners with explanations on what some go functions do i.e goroutines and structs**

**To people who want to contribute: If you would like to contribute to HoneyGO or HoneyComb you'll need to create a fork and submit a Pull Request (PR), The code will be reviewed and potentially be merged (no guarantees). I will also add your name to the contributors list**

As of now we only have basic features to start off with but as development continues we will have more features along with a custom plugin system (HoneyComb, will be available later in the future) and perhaps a way to get forge support although that would be WAYY in the future and I'm not too sure how that would be implemented especially with licensing. Another thing that I want to work on is an easy way to link multiple HoneyGO servers together much like waterfall/bungee, this would be called HoneyPot.

I am very supportive of open-source and free software, for that reason this project will be free forever. You can copy, modify and redistribute my code as long it is under the same license and credit is given to proper authors. If you would like to donate to support HoneyGO that would be much appreciated and you will be added to the donator list (Donators.md)

Checklist:

* [x] 1.15.2 protocol support (Handshake, Status, Login, Play)
* [ ] Entities and Tile entities
* [ ] Mobs and Animals
* [ ] Simple Terrain Generation
* [ ] Chunk/Region Handling
* [ ] Configurable Terrain Generation
* [ ] HoneyGO Plugin API (HoneyComb)
* [ ] Commands
* [ ] Block/Entity Data
* [ ] Actually logging in and being able to play to some degree
* [ ] Bedrock Support
* [ ] Server Configuration
* [ ] HoneyPot (A proxy linking multiple HoneyGO servers muchlike waterfall/bungee)

**Suggestions**

If HoneyGO isn't what you needed you could look at some other projects that are popular and are recommended from me. The projects listed are actively developed and support modern-ish versions of Minecraft (1.12 - 1.15):

Want a server with mods and plugins? Visit [Magma](https://magmafoundation.org/)

Wanting a fast java based Minecraft server? Visit [Paper](https://papermc.io/)

Wanting a stable Java based Minecraft server? Visit [Spigot](https://www.spigotmc.org/)

Wanting a fast, non-mojang based code, server? Visit [Glowstone](https://glowstone.net/)

Wanting a fast Minecraft 1.8 - 1.12.2 server in C/C++? Visit [Cuberite](https://cuberite.org/)

Wanting a Minecraft server with a stable, high level API? Visit [Prismarine](http://flying-squid.prismarine.js.org/#/)

Wanting a paper server with some added gameplay features ? Visit [Purpur](https://github.com/pl3xgaming/Purpur)

Wanting a high performance and highly efficient Minecraft server? Visit [Basin](https://github.com/basinserver/basin)

Wanting an experimental Minecraft Server implementation in Rust? Visit [Feather](https://github.com/feather-rs/feather)

Wanting a Minecraft server that supports the SpongeAPI? Visit [Sponge](https://www.spongepowered.org/) and [Lantern](https://github.com/LanternPowered/Lantern)

<<<<<<< HEAD
**For General use I suggest Paper it is stable, bukkit compatible, fast and is the go to standard for most servers.**
=======
**For General use I suggest Paper it is stable, bukkit compatible, fast and is the go to standard for most servers.** 
>>>>>>> 0d4724463310af3c26cc02324a46d85bbf748cbd

**\
HoneyGO is still an experimental Minecraft Server (Only started development a few months ago) and with that could come major changes that may break/corrupt worlds (We'll try not to). HoneyGO at the moment is only suitable for a simple creative/lobby server as the terrain generation and entity AI is nowhere near completion and could take some time.**
