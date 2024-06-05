PATH=/Users/mdbraber/src/go/bin:~/go/bin/:$PATH
#PATH=~/go/bin/:$PATH
binary=rclone
FRAMEWORK_DIR=build/Release-iphoneos/$binary.framework

gomobile bind -v -target=ios .

xcrun -sdk iphoneos clang -arch arm64 -fpic -shared -Wl,-all_load Rclone.xcframework/ios-arm64/Rclone.framework/Rclone -framework Corefoundation -framework /Users/mdbraber/src/ios_system/build/Release-iphoneos/ios_system.framework/ios_system -lresolv -o $binary.dylib
#ld -syslibroot /Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/Developer/SDKs/iPhoneOS.sdk/ -dynamic -dylib -arch arm64 -lSystem -framework /Users/mdbraber/src/ios_system/build/Release-iphoneos/ios_system.framework/ios_system -dylib_install_name $binary.dylib -all_load -headerpad_max_install_names -weak_reference_mismatches non-weak -o mylib.dylib -platform_version iOS 15 15 Hellogog.xcframework/ios-arm64/Hellogo.framework/Hellogo

cp $binary.dylib ${FRAMEWORK_DIR}/$binary
install_name_tool -id @rpath/$binary.framework/$binary   ${FRAMEWORK_DIR}/$binary

cp ${FRAMEWORK_DIR}/$binary /Users/mdbraber/src/a-shell/$binary.framework/$binary
