#
# spec file for package jenkobs
#
# Copyright (c) 2020 Elektrobit Automotive, Erlangen, Germany.
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

Name:           jenkobs
Version:        0.1
Release:        0
Summary:        Jenkins to OBS connector over AMQP protocol
License:        MIT
Group:          Automotive/Tools
Url:            https://gitlab.com/isbm/jenkobs
Source:         %{name}-%{version}.tar.gz
Source1:        vendor.tar.gz

BuildRequires:  golang-packaging
BuildRequires:  golang(API) >= 1.13

%description
Listener to a RabbitMQ bus via AMQP protocol, used to trigger jobs on Jenkins, based on emitted events from OBS.

%prep
%setup -q
%setup -q -T -D -a 1

%build
# Output for log purpuses
go env

# Build the binary
go build -x -mod=vendor -buildmode=pie -o %{name} ./cmd/jenkobs.go

%install
install -D -m 0755 %{name} %{buildroot}%{_bindir}/%{name}
mkdir -p %{buildroot}%{_sysconfdir}
install -m 0644 ./cmd/jenkobs.conf.example %{buildroot}%{_sysconfdir}/jenkobs.conf
install -m 0644 ./cmd/actions.conf %{buildroot}%{_sysconfdir}/actions.conf

%files
%defattr(-,root,root)
%{_bindir}/%{name}
%dir %{_sysconfdir}
%config /etc/jenkobs.conf
%config /etc/actions.conf

%changelog
