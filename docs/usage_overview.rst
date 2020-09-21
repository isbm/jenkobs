Overview
========

What This Is?
-------------

The **jenkobs** is called after `Jenkins source automation server
<http://jenkins.io>`_ and `OBS (Open Build Service)
<http://openbuildservice.org>`_. It is a daemon for building
integration chains, those usually found in pipelines, continuous
integration etc. The daemon is designed to listen to events
from OBS via message bus and calls specific actions once *something*
is happening with a specific package or project. Depending on demand,
various message events can trigger different actions.

Currently supported actions are:

* **HTTP Action**. It allows ``jenkobs`` to call any URLs via HTTP or
  HTTPS with pre-defined parameters.

* **Shell Action**. Allows ``jenkobs`` to call any shell command with
  specific arguments.


How Does It Work?
-----------------

Generally, OBS needs to run a messaging bus with AMQP protocol. Best
fit in this scenario is to setup `RabbitMQ
<https://www.rabbitmq.com>`_ message broker with OBS as `described in
the documentation
<https://openbuildservice.org/help/manuals/obs-admin-guide/obs.cha.administration.html#_message_bus>`_.

Once AMQP is ready, ``jenkobs`` should have actions configured and
listening to the AMQP. If action criteria matches, then the action is
called and this way other software component is triggered.
