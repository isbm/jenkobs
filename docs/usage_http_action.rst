HTTP Action
===========

Description
-----------

HTTP Action allows calling URLs, accessible via HTTP protocol. It
supports SSL (HTTPS) as well.


Configuration
-------------

Action is configured the following way:

.. code-block:: yaml

   - <PROJECT>:
     status: <AMQP_STATUS>
     package: <PACKAGE>
     arch: <ARCH>
     action:
       type: http
       query:
         url: <PATH> or <URL>
	 method: <METHOD> # post or get, default: get
	 params:
	   <KEY>: <VALUE>

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

* ``METHOD`` is HTTP method, either ``get`` or ``post``. Default is ``get``.

* ``KEY: VALUE`` parameters are arbitrary string/value key pairs. They
  will be included into the HTTP request via ``post`` or ``get``
  methods accordingly.

Values accepting four parameters that are converted then to a values:

1. ``project`` is a Project Name that fired the event.

2. ``package`` is a Package Name, which belongs to the project.

3. ``arch`` is architecture of the Package.

4. ``repo`` is a repository name where project is.
  
Examples
--------

Assuming ``amqp_namespace`` is setup to ``opensuse.obs``, this will
watch all projects when build succeeded and call your localhost at
8080 port, sending project name:

.. code-block:: yaml

   - "*":
     status: opensuse.obs.package.build_success
     action:
       type: http
       query:
         url: http://localhost:8080
	 params:
	   name: "{project}"

With the same configuration assumptions as above, this action will
watch ``libsolv`` package inside ``YaST:Head`` on Open Build Service
and will trigger configured Jenkins job, called ``libsolv``:

.. code-block:: yaml

   - "YaST:Head":
     status: opensuse.obs.package.build_success
     package: libsolv
     arch: x86_64
     action:
       type: http
       query:
         url: /job/libsolv/build
