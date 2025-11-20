package WorkSamples

check.Progress(ctx, "Checking if crowdstrike service is running")

conn, err := dbus.NewSystemdConnectionContext(ctx)
if err != nil {
log.ErrorContextf(ctx, "Failed to create dbus connection: %v", err)
return result.Failf("Failed to create dbus connection").WithCode(22), nil
}
defer conn.Close()

state, err := conn.ListUnitsByNamesContext(ctx, []string{falconService})
if err != nil {
log.ErrorContextf(ctx, "Failed to list units: %v", err)
return result.Failf("Failed to list units").WithCode(13), nil
}
if strings.TrimSpace(state[0].ActiveState) != "active" {
log.ErrorContextf(ctx, "falcon-sensor.service  is not running: %v", state[0].ActiveState)
return result.Failf("falcon-sensor.service is not running").WithCode(18), startCrowdStrike
}

return result.Pass("CrowdStrike service seems to be in a healthy state").WithCode(21), nil
}