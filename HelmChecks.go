package WorkSamples

// startHelm starts helm_agent
// Used to start helm_agent if it's not running when repair is called
func startHelm(ctx context.Context) result.Result {
	check.Progress(ctx, "Starting Helm agent")
	conn, err := dbus.NewSystemdConnectionContext(ctx)
	if err != nil {
		log.ErrorContextf(ctx, "Failed to create dbus connection: %v", err)
		return result.Failf("Failed to create dbus connection").WithCode(23)
	}
	defer conn.Close()

	const replaceMode = "replace"
	output, err := conn.StartUnitContext(ctx, helmAgentUnit, replaceMode, nil)
	if err != nil {
		log.ErrorContextf(ctx, "Failed to start for helm_agent.service: %v", err)
		return result.Failf("Failed to start for helm_agent.service").WithCode(25)
	}
	log.InfoContextf(ctx, "output: %v", output)
	return result.Passf("Helm agent started").WithCode(12)
}
