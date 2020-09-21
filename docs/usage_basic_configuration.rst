Basic Configuration
===================

The ``jenkobs`` daemon has two configuration files:

- Basic configuration, usually called ``jenkobs.conf``.
- Actions configuration, specified in the basic configuration.

Configuration files by default are searched in the following order:

1. ``jenkobs.conf`` in the same directory as the binary
2. ``~/.jenkobs`` in user HOME
3. ``/etc/jenkobs.conf`` in the global ``/etc`` directory

Setup
-----

There are two sections currently:

- ``jenkins``. This section is specifically for Jenkins connection
- ``amqp``. Section, where AMQP parameters are defined

Additionally, there is one key ``actions`` which is a full path to the
actions configuration file.

Example
-------

Below is an example of a configuration:

.. code-block:: yaml

   jenkins:
     username: john
     token: 1234567890
     hostname: yourjenkins.com
     port: 8080

   amqp:
     username: opensuse
     password: opensuse
     fqdn: amqp.opensuse.org
     port: 0 # Disable port or specify one explicitly
     exchange: pubsub
     vhost:  # Declare one, if you have it

   actions: /path/to/your/actions.conf

     
