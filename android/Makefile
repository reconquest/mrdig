ANDROID_PACKAGE=io.reconquest.mrdig

# Settings for keystore which is used to sign APK.
KEYS_DN        = DC=io,CN=reconquest
KEYS_PASS      = 123456
KEYS_VALIDITY  = 365
KEYS_ALGORITHM = RSA
KEYS_SIZE      = 2048
KEYS_ALIAS     = mrdig

_ADB=adb -s $(shell adb devices -l | tail -n+2 | cut -f1 -d' ' | head -n1)

ifdef FASTBUILD
GRADLE_BUILD_FLAGS = -x lint -x lintVitalRelease
else
GRADLE_BUILD_FLAGS =
endif

run: install
	$(_ADB) shell am start -n $(ANDROID_PACKAGE)/.MainActivity

install: build/mrdig.apk
	$(_ADB) install -r build/mrdig.apk

.keystore:
	@echo :: initializing keystore
	@keytool -genkeypair \
		-alias $(KEYS_ALIAS) \
		-validity $(KEYS_VALIDITY) \
		-keystore $@ \
		-keyalg $(KEYS_ALGORITHM) \
		-keysize $(KEYS_SIZE) \
		-storepass $(KEYS_PASS) \
		-keypass $(KEYS_PASS) \
		-dname $(KEYS_DN) \
		-deststoretype pkcs12

build/mrdig.apk: .keystore java

java: .keystore
	gradle assembleDebug $(GRADLE_BUILD_FLAGS)
	mv build/outputs/apk/debug/android-debug.apk build/mrdig.apk

clean:
	rm -rf build
	go clean -cache

eclipse:
	gradle eclipse
	gradle pom

eclipse-clean:
	rm -rf .classpath .project .settings
