#! /bin/bash
rsync -zrvh -e ssh export/* zito@demonzito.com:/home/zito/dating.demonzito.com
rsync -zrvh -e ssh assets zito@demonzito.com:/home/zito/dating.demonzito.com