### OkSetda Absensi
OkSetda e-Attendance Telegram Long-Polling Bot

#### Deployment Instructions

Follow this instruction below to deploy this application on the server at no pain. But, you can run via Docker & docker-compose if you didn't have a time, KISS (KEEP IT SIMPLE STUPID) #LOL!

##### System Requirement
1. Basic/Standard Server
2. Database (same as server_platform)
3. GOLANG version ^1.16
4. Docker & docker-compose (if you want to use docker) *(optional)

##### Deploy via Docker & docker-compose
1. Copy `.env.example` file
    ```shell
    $ cp .env.example .env
    ```
2. Update Environment Variable [Database Connection, Credentials, etc.]
3. Run this app as container with docker-compose
    ```shell
    $ docker-compose up -d
    ```
   
#### Deploy and run via Systemd
1. Copy `.env.example` file
    ```shell
    $ cp .env.example bin/.env
    ```
2. Update Environment Variable [Database Connection, Credentials, etc.]
3. Build Binaries
   ```shell
   $ go build -o bin/svc-tgbotgo
   ```
4. Update systemd configuration file at `bin/systemd/svc-tgbotgo.service`
   - path_to_service `$ pwd` <- use this command to get current project directory
   - selected_user 
   - selected_group 
5. copy systemd configuration file
   - from `bin/systemd/svc-tgbotgo.service`
   - to `/lib/systemd/system/`
   ```shell
   $ sudo cp bin/systemd/svc-tgbotgo.service /lib/systemd/system/
   ```
6. grant permission to systemd service 
   ```shell
    $ sudo chmod 755 /lib/systemd/system/svc-tgbotgo.service
   ```
7. Now, you should be able to enable the service, start it, then monitor the logs by tailing the systemd journal:
   ```shell
    $ sudo systemctl enable svc-tgbotgo.service

    $ sudo systemctl start svc-tgbotgo.service
   ```
8. Querying and displaying logs from journald, systemd's logging service.
   ```shell
    $ sudo journalctl -f -u sleepservice
   ```
9. for more information see [running-a-go-binary-as-a-systemd-service](https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/)