# Freon Service [![Build Status](https://travis-ci.org/Joystream/substrate-runtime-joystream.svg?branch=master)](https://travis-ci.org/Joystream/substrate-runtime-joystream)

The project is a ready-made box solution for comfortable work with multilingual applications. 

You can generate translations for different platforms (web services, mobile applications, third-party grpc microservices) in the same style inside the admin panel of this service, and also control the work of your translators.

This service is an embedded solution into your infrastructure to increase flexibility and reliability. Thanks to the docker build, you can easily add this service to your infrastructure.
The service consists of two parts - a backend written in Golang and a frontend written in ReactJS.

This service will allow content managers to conveniently make changes to localization files, monitor support for various languages, and seamlessly pull them up for applications without the need to update in the application store (App Store, Google Play, etc.) or instant update in web app.

## Build Status

Development, beta-release

### Building

The Freon service consists of two parts - backend (Golang) and frontend (ReactJS). The Web Interface is compiled as a ready-made bundle with the assembled JS code, which is launched through the proxy server built into the backend for quick and easy access to the admin panel.

All information is stored in Postgres database, static documents (generated translations) are stored in volume docker image and can be easily requested as static files through URL.

Communication interface with Freon service:
- **Rest API.** Used to interact with mobile applications, services and the admin panel.
- **GRPC.** Used to interact with microservices inside your application.

The following ports are specified by default:
- **4000** - Rest API
- **4001** - GRPC
- **4002** - Proxy server for working with the admin panel.

When you start the Freon service for the first time, a default administrator will be created, in the future you can delete it or change its username and password.
- **Login:** admin@admin.com
- **Password:** adminpassword

**ATTENTION!** Don't use this credential in production mode. Your data could be stolen.  

#### Docker quick start

One of the quickest ways to get Freon up and running on your machine is by using Docker:

```shell
docker run -d --name freon -v freon-docs:/docs \
           --env DB_HOST=localhost --env DB_PORT=5432 \
           --env DB_USER=root --env DB_PASS=password \
           --env DB_NAME=freon --env JWT_SECRET_KEY=freon-secret-key \
           -p 4000:4000 -p 4001:4001 -p 4002:4002 \
           freonservice/freonservice
```

Volume `freon-docs:/docs` use as default place for storing generated translations for web, android and ios applications.

For quick launch and deployment you can use `docker-compose.prod.yml` file with already configured parameters for this service.

Command for launch service from docker

`docker-compose -f docker-compose.prod.yml up -d`

## Example of interface

![Translation page](https://github.com/freonservice/freon/blob/master/freon-page.png)

## Future features

- [ ] Connection google-translation api for quick semi-automation translations
- [ ] Generation i18n files for web
- [ ] Generation android translations
- [ ] Generation ios translations
- [ ] UI improvements
- [x] Add API and business logic for checking actual version of translations  
- [ ] Fix supporting pluralization
- [ ] Add AWS storages for static files

## Contribution

Thank you for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes!

If you liked this project, we will be happy to welcome you to our team.
To maintain a consistent style code, we ask you to use our *linter* and *tests* before each pull requests.

You can install *git hooks* for convenient work:

`./scripts/install-hooks.bash`

We try to keep the project clean, so we write tests and use the linter.

## License

This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details