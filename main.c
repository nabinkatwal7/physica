#include <stdio.h>
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xos.h>

void init_x(){
	unsigned long black, white;

	dis = XOpenDisplay(NULL);
}

int main()
{
	init_x();
	return 0;
}
