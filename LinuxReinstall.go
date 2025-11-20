package WorkSamples

func main() {

	ctx := context.Background()

	fmt.Println("Welcome to the gLinux Reimage script!\nGive the script a few seconds to run through system pre-checks.")

	fmt.Println("Note: Press Ctrl + C to exit at any time")

	if err := needsA11y(); err != nil {
		log.ExitContextf(ctx, "Not proceeding: user requires a11y tools support")
	}
	role, err := currentRole()
	if err != nil {
		log.ExitContextf(ctx, "Failed at isRole: %v", err)
	}
	if err := ensureSupportedDevice(); err != nil {
		log.ExitContextf(ctx, "Failed at ensureSupportedDevice: %v", err)
	}
	if err := ensureSingleDisk(); err != nil {
		log.ExitContextf(ctx, "Failed at ensureSingleDisk: %v", err)
	}
	if err := secureBootEnabled(ctx); err != nil {
		log.ExitContextf(ctx, "Failed at Secure Boot Check: %v", err)
	}
	if err := ensureNoMultiUser(ctx); err != nil {
		log.ExitContextf(ctx, "Failed at ensureNoMultiUser: %v", err)
	}
	if err := checkInventoryData(ctx); err != nil {
		log.ExitContextf(ctx, "Failed at Inventory Data Check: %v, \nContact go/techstop to get this error corrected", err)
	}
	if err := biosUpToDate(ctx); err != nil {
		log.ExitContextf(ctx, "Failed at BIOS Check: %v", err)
	}
	fmt.Println("-----------------------------------------------------------------------")
	if err := onLegalHold(); err != nil {
		log.ExitContextf(ctx, "Not proceeding: user on legal hold")
	}
	fmt.Println("-----------------------------------------------------------------------")
	if err := stablePower(role); err != nil {
		log.ExitContextf(ctx, "Failed at Stable Power Check: %v", err)
	}
	if err := homeDirRun(ctx); err != nil {
		log.ExitContextf(ctx, "Failed at HomeDirHelper: %v", err)
	}
	fmt.Println("-----------------------------------------------------------------------")
	if err := runBackups(ctx); err != nil {
		log.WarningContextf(ctx, "Failed at Backups: %v", err)
	}
	fmt.Println("-----------------------------------------------------------------------")
	if err := checkRecoveryInstaller(ctx); err != nil {
		log.WarningContextf(ctx, "Failed at Recovery Installer Check: %v", err)
	}
	if err := maskShutdownInstaller(); err != nil {
		log.WarningContextf(ctx, "Failed at Masking Shutdown Installer: %v", err)
	}
	if err := saveWifiProfile(ctx); err != nil {
		log.WarningContextf(ctx, "Failed at Saving Wifi Profile: %v", err)
	} // If it works or not has not consequence for the reimage > just won't be touchless.
	fmt.Println("-----------------------------------------------------------------------")
	role, err = maybeSwitchRole(role)
	if err != nil {
		log.ExitContextf(ctx, "Failed at maybeSwitchRole: %v", err)
	}
	fmt.Println("-----------------------------------------------------------------------")
	if !prompt.Confirm("that you would like to continue with reimaging", "Continuing with reimaging will erase user data and reboot the system now", false) {
		log.ExitContextf(ctx, "Exiting...")
	}
	if err := reimage(ctx, role); err != nil {
		log.ExitContextf(ctx, "Failed at Reimage: %v", err)
	}
}
