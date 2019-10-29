ANDROID_PACKAGE=io.reconquest.mrdig

run-%:
	$(MAKE) ebiten-$*
	$(MAKE) -C $* ANDROID_PACKAGE=$(ANDROID_PACKAGE) run

ebiten-android:
	GO111MODULE=off \
		ebitenmobile bind \
		-target android \
		-javapkg $(ANDROID_PACKAGE) \
		-o android/build/game.aar \
		mrdig/game
