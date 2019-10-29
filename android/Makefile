ANDROID_PACKAGE=io.reconquest.mrdig

# Settings for keystore which is used to sign APK.
KEYS_DN        = DC=io,CN=reconquest
KEYS_PASS      = 123456
KEYS_VALIDITY  = 365
KEYS_ALGORITHM = RSA
KEYS_SIZE      = 2048
KEYS_ALIAS     = mrdig

_ADB=adb -s $(shell adb devices -l | tail -n+2 | cut -f1 -d' ' | head -n1)

_GO_SRC=$(shell find ../game/ -type f)

_flags: $(filter +%,$(MAKECMDGOALS))

+fast: +nolint +noaar

+nolint:
	$(eval GRADLE_BUILD_FLAGS += -x lint -x lintVitalRelease)

+noaar:
	$(eval _AAR=:)

debug@run: debug@install
	$(_ADB) shell am start -n $(ANDROID_PACKAGE)/.MainActivity

debug@install: debug@apk
	$(_ADB) install -r build/mrdig.apk

build/game.aar: $(_GO_SRC)
	@$(_AAR) echo :: building game aar
	@$(_AAR) GO111MODULE=off \
		ebitenmobile bind \
		-target android \
		-javapkg $(ANDROID_PACKAGE) \
		-o build/game.aar \
		github.com/reconquest/mrdig/game

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

debug@apk: _flags .keystore build/game.aar
	@echo :: building apk
	@gradle assembleDebug $(GRADLE_BUILD_FLAGS)
	@mv build/outputs/apk/debug/android-debug.apk build/mrdig.apk

debug: debug@run

clean:
	@rm -rf build
	@go clean -cache

eclipse:
	@gradle eclipse
	@gradle pom

eclipse-clean:
	@rm -rf .classpath .project .settings
