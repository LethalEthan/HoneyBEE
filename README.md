# HoneyBEE

HoneyBEE is a Minecraft server written in GO, with help from [wiki.vg](https://wiki.vg) (Huge thanks to them and I suggest helping them out with a donation for the awesome resources that they publish for free <3).

Join my discord for support and to ask questions :) [here](https://discord.gg/rBm2U5TJNx)

This software is still being developed so expect bugs and issues now and then, also remember that feedback and disclosure of issues is welcomed but please don't send any hateful or spam issues :). HoneyBEE as of now is considered pre-alpha it is very early in development and has only been released publicly on github since 25/4/2020

**Disclaimer: All code is completely original and is not based off of Mojang's code, any PR's must not be based off of or have any of Mojang's copyrighted code as it will not be accepted nor be encouraged as it would be obviously illegal, please see the Terms and Conditions of Mojang for familiarisation. I am not in any way affiliated with Mojang, HoneyBEE is not endorsed nor approved by Mojang.** 

**This project is a cleanroom project meaning that I have to reverse engineer and read documentation that others and have wrote to understand how the network protocol works along with any other kind of necessary data, world generation will be completely be based off my own work along with AI and other ideas**

**To people who want to contribute: If you would like to contribute to HoneyBEE or HoneyComb you'll need to create a fork and submit a Pull Request (PR), The code will be reviewed and potentially be merged (no guarantees). I will also add your name to the contributors list. Any PR'sand commits must use the same license and a CLA must be signed**

As of now we only have basic features to start off with but as development continues we will have more features along with a custom plugin system (HoneyComb, will be available later in the future) and perhaps a way to get forge support although that would be WAYY in the future and I'm not too sure how that would be implemented especially with licensing. 

I am very supportive of open-source and free software, for that reason this project will be free forever for **personal use only**. You can copy, modify and redistribute my code as long it is under the terms of the license and credit is given to proper authors. If you would like to donate to support HoneyBEE that would be much appreciated and you will be added to the donator list (Donators.md)

I would like to give a massive thanks to: @lynxplay @MiniDigger @magiccheese1 @mrsherobrine @Felenov @fahlur and @SNokerYT

# Checklist

* [x] 1.17 protocol support (Handshake, Status, Login, Play)
* [ ] Entities and Tile entities
* [ ] Mobs and Animals
* [ ] Simple Terrain Generation
* [ ] Chunk/Region Handling
* [ ] Configurable Terrain Generation
* [ ] HoneyBEE Plugin API (HoneyComb)
* [ ] Commands
* [ ] Block/Entity Data
* [X] Actually logging in and being able to play to some degree
* [X] NBT
* [ ] Bedrock Support
* [x] Server Configuration
* [ ] HoneyHive (A proxy linking multiple HoneyBEE servers muchlike waterfall/bungee)

# Suggestions

**For General use, I suggest Paper it is stable, bukkit compatible, fast and is the go to standard for most servers.**

**HoneyBEE is in pre-alpha and with that could come major changes that may break/corrupt worlds (We'll try not to). HoneyBEE at the moment will only suitable for a simple creative/lobby server as the terrain generation and entity AI is nowhere near completion and could take some time.**

If HoneyBEE isn't what you needed you could look at some other projects that are popular and are recommended from me. The projects listed are actively developed and support modern-ish versions of Minecraft (1.12 - 1.17):

Wanting a lightweight C# minecraft server? [Starfield](https://github.com/StarfieldMC/Starfield) (C#)

Wanting a Java based Minecraft server? Visit [Spigot](https://www.spigotmc.org/) (Java)

Wanting a fast java based Minecraft server? Visit [Paper](https://papermc.io/) (Java)

Wanting a scalable, experimental server in C#? visit [Mincase](https://github.com/dotnetGame/MineCase) (C#)

Wanting a server with forge mods and bukkit plugins? Visit [Magma](https://magmafoundation.org/) (Java)

Wanting a fast Minecraft 1.8 - 1.12.2 server in C/C++? Visit [Cuberite](https://cuberite.org/) (C++)

Wanting a multithreaded Minecraft server built for redstone? [MCHPR](https://github.com/MCHPR/MCHPRS) (Rust)

Wanting a Minecraft server with a stable, high level API? Visit [Prismarine](http://flying-squid.prismarine.js.org/#/) (Js)

Wanting a C# implementation of the Minecraft server protocol. [Obsidian](https://github.com/ObsidianMC/Obsidian) (C#)

Wanting a fast, lightweight Minecraft server written in Kotlin? [Krypton](https://github.com/KryptonMC/Krypton) (Kotlin)

Wanting a paper server with some added gameplay features ? Visit [Purpur](https://github.com/pl3xgaming/Purpur) (Java)

Want to create your own server with an API other than bukkit? Visit [Minestom](https://github.com/Minestom/Minestom) (Java)

Wanting an experimental Minecraft Server implementation in Rust? Visit [Feather](https://github.com/feather-rs/feather) (Rust)

Wanting a Minecraft server that supports the SpongeAPI? Visit [Sponge](https://www.spongepowered.org/) and [Lantern](https://github.com/LanternPowered/Lantern) (Java)

Wanting a fast, cleanroom based server implementation with bukkit? Visit [Glowstone](https://glowstone.net/) (Java)

Wanting a fast multithreaded and async implementation of a minecraft server? [MotorMC](https://github.com/garet90/MotorMC) (C)

Wanting a high performance minecraft server written in golang, soley made for PvP? [gogs](https://github.com/GambitLLC/gogs) (GO)

Wanting a high performance paper-fork server with additional performance patches? [airplane](https://airplane.gg/) (Java)
