+%: # accept flags for nested Makefiles
debug~%:
	@$(MAKE) --no-print-directory -s -C $* debug $(filter +%,$(MAKECMDGOALS))
