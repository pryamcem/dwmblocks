/* modify this file to change what commands output to your statusbar, and recompile using the make command. */

static const Block blocks[] = {
    /* Icon         Command                                                                     Interval    Signal */
    { "",           "/home/pryamcem/.config/dwmblocks/scripts/sb-ram2",                             2,         9       },
    { "",           "/home/pryamcem/.config/dwmblocks/scripts/sb-volume2",                          0,         10      },
    { "",           "/home/pryamcem/.config/dwmblocks/scripts/sb-battery",                         30,         8       },
    //Yaremche
    //{ "",           "/home/pryamcem/.config/dwmblocks/scripts/weather/goweather 48.4372 24.5694", 600,         11       },
    //Lviv
    { "",           "/home/pryamcem/.config/dwmblocks/scripts/weather/goweather 49.8383 24.0232", 600,         11       },
    { "",           "/home/pryamcem/.config/dwmblocks/scripts/moon/gomoon",                      3600,         12       },
    { "",           "/home/pryamcem/.config/dwmblocks/scripts/sb-datetime",                        10,         7       },
};

/* sets delimeter between status commands. NULL character ('\0') means no delimeter. */
static char delim[] = "  ";
static unsigned int delimLen = 4;
