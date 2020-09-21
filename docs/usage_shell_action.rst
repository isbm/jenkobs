Shell Action
===========

Description
-----------

Shell Action allows calling arbitrary shell commands if a specific
project is triggered. This is useful mostly for setting up environment
via Configuration Management tools (e.g. Ansible), logging etc.


Configuration
-------------

Action is configured the following way:

.. code-block:: yaml

   - <PROJECT>:
     status: <AMQP_STATUS>
     package: <PACKAGE>
     arch: <ARCH>
     action:
       type: shell
       command:
         - <NAME>
	 - <OPTION>
	 - <OPTION>
	 - ...

* ``PROJECT`` (mandatory) can be a specific name of a project in OBS, or an
  asterisk ``*`` to match everything at once.

* ``AMQP_STATUS`` (mandatory) is configured action. For example, according to the
  `OBS Admin Guide
  <https://openbuildservice.org/help/manuals/obs-admin-guide/obs.cha.administration.html#idm140614333062832>`_,
  it is ``amqp_namespace`` + ``object.action``, as described on `AMQP
  server of OpenSUSE package building website <https://amqp.opensuse.org>`_.

* ``PACKAGE`` (optional) is a package name to watch. This will limit
  watcher only to a specific package in a specific project. If
  ``PACKAGE`` option is not defined, any package within the project
  firing an event will trigger the action.

* ``ARCH`` (optional) is an architecture name, such as ``x86_64`` or
  ``armv8`` etc.
  
* ``PATH`` or ``URL``. If ``PATH`` (mandatory) is defined, then it will be called
  on the  the current configured Jenkins (see the basic
  configuration). Otherwise the ``URL`` will be called directly,
  omitting any Jenkins configuration specifics.

* Essentially a command is as those one would typically construct with
  the CLI, something like ``name -option`` etc. The interface is quite
  simple: whenever you need a space, it should be a next option in the
  YAML configuration.

Options accepting four parameters that are converted then to a values:

1. ``project`` is a Project Name that fired the event.

2. ``package`` is a Package Name, which belongs to the project.

3. ``arch`` is architecture of the Package.

4. ``repo`` is a repository name where project is.

Examples
--------

Assuming ``amqp_namespace`` is setup to ``opensuse.obs``, this will
watch all projects when build succeeded and then add a message to a
system logger:

.. code-block:: yaml

   - "*":
     status: opensuse.obs.package.build_success
     action:
       type: shell
       command:
         - logger
	 - "Project {project} successfully published to {repo} repository"

