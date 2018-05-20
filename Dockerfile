FROM cloudwarelabs/xorg:latest
RUN apt-get update
RUN apt-get install dictionaries-common
RUN /usr/share/debconf/fix_db.pl && dpkg-reconfigure dictionaries-common
RUN apt-get install -y gnome-themes-standard xfce4
RUN apt-get remove -y xscreensaver xscreensaver-data
RUN mkdir -p /root/.config/xfce4/xfconf/xfce-perchannel-xml
COPY xsettings.xml /root/.config/xfce4/xfconf/xfce-perchannel-xml/
COPY xfce4-panel.xml /root/.config/xfce4/xfconf/xfce-perchannel-xml/
RUN mkdir -p /root/.config/autostart

RUN apt-get update
RUN apt-get install -y libwebp-dev libx11-dev libxdamage-dev libxtst-dev libpng12-0
COPY build/pulsar /usr/local/bin/pulsar
COPY libwebsockets.so.11 /usr/lib/
COPY pulsar.desktop /root/.config/autostart/
ENV DISPLAY :0
ENV PULSAR_PORT 5678
EXPOSE 5678

COPY

ENTRYPOINT  ["run.sh"]