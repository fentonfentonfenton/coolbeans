# coolbeans!
Easily connect to an SSM Managed Instance from your terminal.

If you have SSM Managed Instances - and have pretty much disabled `SSH` LIKE A BOSS, then you might want to be able to connect to your instances super easily without faffing around in the AWS UI.

If your EC2 Instances have easily idenfitiable tags - you might be in luck. This will try until one of them lets it connect, then tell you what to do. Or it will exit gracefully.

### Install:

* Install the SSM Plugin via https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html

* Grab the binary (only tested for OSX today but will work on WIN10 and Linux with a light touch) [m1](https://github.com/fentonfentonfenton/coolbeans/releases/download/0.0.1/coolbeans-m1) / [osx](https://github.com/fentonfentonfenton/coolbeans/releases/download/0.0.1/coolbeans-osx)

* `./coolbeans-osx example.com` - where the argument is a tag value you have for an instance

* !!!!?

* Paste the final output string and hit enter - 
```Connecting to AWS
/usr/local/bin/aws ssm start-session --target i-123abc


curl -H "X-aws-ec2-metadata-token: $TOKEN" -v http://169.254.169.254/latest/meta-data/instance-id

HTTP/1.1 200 OK

i-123abc
```

* Profit


### TO:DO

Make this read `stdin` so you can be interactive.
Tests
LINUX/WINDOWS
Use different paramaters to identify servers.


