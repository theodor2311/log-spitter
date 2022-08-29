# log-spitter
This is a testing application written in GO to flush logs for performance testing.

The following environment variable can be added to adjust the log spitting behavior.
| Environment Variable | Description                                   | Default Value |
|----------------------|-----------------------------------------------|---------------|
| LOG_COUNT            | The number of logs sends per period            | 20000         |
| PERIOD               | The period of time in ms between sending logs | 0             |
| MAX_COUNT            | The total of logs will be sent                | 1000000       |

With the default value as period=0, the logs will be flushed non-stop until 1000000 has been flushed.
If the period is set to 1000, after sending 20000 logs it will pause for 1 second before continuing sending the logs, this can help to suppress the logs per second sent by this testing application.

My built version can be pulled from docker.io/theodor2311/log-spitter:latest.
