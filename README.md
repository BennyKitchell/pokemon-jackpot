<a id="readme-top"></a>
# Pokemon Jackpot
It's Pokemon, it's a slot machine, it's silly fun


<!-- ABOUT THE PROJECT -->
## About The Project

Languages used: Go, React, Typescript

This project was created for a tak-home assignment by Rockbot. It was an exercise to work on a larger project using Go. 

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started
### Prerequisites


Be sure to have the following installed:
* Go 1.23.0
    * brew
    ```sh
        brew install go
    ```
* Docker
    * brew
    ```sh
        brew install docker
    ```

* Vite
     * brew
    ```sh
        brew install vite
    ```



### Installation & Usage

_After cloning the package run the following to get started_

1. Clone the repo
   ```sh
   git clone https://github.com/BennyKitchell/pokemon-jackpot.git
   ```
2. In the root of the project, run `./db-init.sh` 
    _This will start Kafka, Redis, and the DB then populate the db with pokemon_
3. ```sh
        cd user-service/cmd
        go run .
    ```
    _This will start the user service_
4. In a new terminal window run 
    ```sh
        cd pokemon-service/cmd
        go run .
    ```
    _This will start the pokemon service_

4. In a new terminal window run 
    ```sh
        cd client/
        npm install
        vite
    ```
    _This will start the client at localhost:5173_
5. In a browser window visit `localhost:5173`

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [ ] More tests with more in depth tests
- [ ] Add Changelog
- [ ] Dockerize all of the services
- [ ] Better UI with notifications about new pokemon and user already exists warnings
- [ ] Deploy with Kubernetes for better scaling

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
