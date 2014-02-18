package main
/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AVFoundation -framework CoreMedia

#import <AVFoundation/AVFoundation.h>

static int _playURL(const char* url) {
    @autoreleasepool {
        NSURL* u = [NSURL URLWithString:[NSString stringWithUTF8String:url]];

        AVPlayer* p = [AVPlayer playerWithURL:u];
        [p play];

        NSTimeInterval played = 0.;
        while (1) {
            NSTimeInterval t = CMTimeGetSeconds([p currentTime]);
            if (t > 0. && t == played) {
                break;
            }
            played = t;

            [[NSRunLoop currentRunLoop] runMode:NSDefaultRunLoopMode
                                     beforeDate:[[NSDate date] dateByAddingTimeInterval:.1]];

        }

        return 0;
    }
}

*/
import "C"
import "unsafe"

import (
	"errors"
)

func playURL(url string) error {
	c := C.CString(url)
	defer C.free(unsafe.Pointer(c))

	if r := C._playURL(c); r != 0 {
		return errors.New("play error")
	}
	return nil
}

