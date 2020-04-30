package main

import (
    "flag"
    "log"
	"fmt"
    "strings"
    "github.com/prometheus/client_golang/prometheus"
	nut "github.com/robbiet480/go.nut"
)

const (
    namespace = "nutexporter"
)


var (
        telemetryAddr = flag.String("telemetry.addr", ":9162", "address for nut exporter")
        metricsPath   = flag.String("telemetry.path", "/metrics", "URL path for surfacing collected metrics")

        nutUpsName    = flag.String("nut.ups_name","","NUT UPS name")
        nutAddr       = flag.String("nut.addr", "", "address of Network UPS tools server")
        nutPort       = flag.String("nut.port", "3493", "port of Network UPS tools server")
        nutUsername   = flag.String("nut.username", "", "NUT username")
        nutPassword   = flag.String("nut.password", "", "NUT password")
)

func main() {

	flag.Parse()
	if *nutUpsName == "" {
	    log.Fatal("NUT UPS name must be specified with '-nut.ups_name' flag")
	}

	if *nutAddr == "" {
	    log.Fatal("NUT Server address must be specified with '-nut.addr' flag")
	}

//        fn := newClient(*apcupsdNetwork, *apcupsdAddr)

//        prometheus.MustRegister(apcupsdexporter.New(fn))

//        http.Handle(*metricsPath, prometheus.Handler())
//        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//                http.Redirect(w, r, *metricsPath, http.StatusMovedPermanently)
//        })

    
    log.Printf("starting nut exporter on %q for server %s://%s",
        *telemetryAddr, *nutUpsName, *nutAddr)

//        if err := http.ListenAndServe(*telemetryAddr, nil); err != nil {
//                log.Fatalf("cannot start apcupsd exporter: %s", err)
//        }


//moje 
    //GetUPSStatus()

}

var _ prometheus.Collector = &UPSCollector{}

func GetUPSStatus() {
    //fmt.Printf("%T\n",*nutAddr)
	client, connectErr := nut.Connect(*nutAddr)
	if connectErr != nil {
        fmt.Println("ERR: %s",*nutAddr)
		fmt.Print(connectErr)

	}
//	_, authenticationError = client.Authenticate(*nutUsername, *nutPassword)
//	if authenticationError != nil {
//		fmt.Print(authenticationError)
//	}
    upsCMD := "LIST VAR " + *nutUpsName
    upsList, listErr := client.SendCommand(upsCMD)
	if listErr != nil {
		fmt.Print(listErr)
	}
    for _, el := range upsList {
        //fmt.Println(strings.Replace(el,"VAR " + *nutUpsName + " ","",1))
        if strings.HasPrefix(el,"VAR") {
            ParseKV(strings.Replace(el,"VAR " + *nutUpsName + " ","",1))
        }
    }
}

func ParseKV(kv string) {
    sp := strings.SplitN(kv, " ",2)
    fmt.Println(sp[0],sp[1])
    return
}

func (s *UPSvals) parseKVString(k string,v string) bool {
    switch k {
    case "battery.charge":
        s.ups_battery_charge = v
    case "battery.voltage":
        s.ups_battery_voltage = v
    case "battery.voltage.high":
        s.ups_battery_voltage_high = v
	case "battery.voltage.low":
        s.ups_battery_voltage_low = v
	case "battery.voltage.nominal":
        s.ups_battery_voltage_nominal = v
	case "device.type":
        s.ups_device_type = v
	case "driver.name":
        s.ups_driver_name = v
	case "driver.parameter.bus":
        s.ups_driver_parameter_bus = v
	case "driver.parameter.pollfreq":
        s.ups_driver_parameter_pollfreq = v
	case "driver.parameter.pollinterval":
        s.ups_driver_parameter_pollinterval = v
	case "driver.parameter.port":
        s.ups_driver_parameter_port = v
    case "driver.parameter.productid":
        s.ups_driver_parameter_productid = v
	case "driver.parameter.synchronous":
        s.ups_driver_parameter_synchronous = v
	case "driver.parameter.vendorid":
        s.ups_driver_parameter_vendorid = v
	case "driver.version":
        s.ups_driver_version = v
	case "driver.version.data":
        s.ups_driver_version_data = v
	case "driver.version.internal":
        s.ups_driver_version_internal = v
	case "input.voltage":
        s.ups_input_voltage = v
	case "input.voltage.fault":
        s.ups_input_voltage_fault = v
	case "output.current.nominal":
        s.ups_output_current_nominal = v
	case "output.frequency":
        s.ups_output_frequency = v
	case "output.frequency.nominal":
        s.ups_output_current_nominal = v
	case "output.voltage":
        s.ups_output_voltage = v
	case "output.voltage.nominal":
        s.ups_output_voltage = v
	case "ups.beeper.status":
        s.ups_ups_beeper_status = v
	case "ups.delay.shutdown":
        s.ups_ups_delay_shutdown = v
	case "ups.delay.start":
        s.ups_ups_delay_start = v
	case "ups.firmware.aux":
        s.ups_ups_firmware_aux = v
	case "ups.load":
        s.ups_ups_load = v
	case "ups.productid":
        s.ups_ups_productid = v
	case "ups.status":
        s.ups_ups_status = v
	case "ups.type":
        s.ups_ups_type = v
    case "ups.vendorid":
        s.ups_ups_vendorid = v
    default:
        return false
    }
    
    return true

}
type UPSvals struct {
	ups_battery_charge string
	ups_battery_voltage string
	ups_battery_voltage_high string
	ups_battery_voltage_low string
	ups_battery_voltage_nominal string
	ups_device_type string
	ups_driver_name string
	ups_driver_parameter_bus string
	ups_driver_parameter_pollfreq string
	ups_driver_parameter_pollinterval string
	ups_driver_parameter_port string
	ups_driver_parameter_productid string
	ups_driver_parameter_synchronous string
	ups_driver_parameter_vendorid string
	ups_driver_version string
	ups_driver_version_data string
	ups_driver_version_internal string
	ups_input_voltage string
	ups_input_voltage_fault string
	ups_output_current_nominal string
	ups_output_frequency string
	ups_output_frequency_nominal string
	ups_output_voltage string
	ups_output_voltage_nominal string
	ups_ups_beeper_status string
	ups_ups_delay_shutdown string
	ups_ups_delay_start string
	ups_ups_firmware_aux string
	ups_ups_load string
	ups_ups_productid string
	ups_ups_status string
	ups_ups_type string
	ups_ups_vendorid string
}

type UPSCollector struct {
    Info                    *prometheus.Desc
    BatteryChargePercent    *prometheus.Desc
    BatteryVoltage          *prometheus.Desc
    BatteryVoltageLow       *prometheus.Desc
    BatteryVoltageHigh      *prometheus.Desc
    BatteryVoltageNominal   *prometheus.Desc
    InputVoltage            *prometheus.Desc
    InputVoltageFault       *prometheus.Desc
    OutputCurrentNominal    *prometheus.Desc
    OutputFrequency         *prometheus.Desc
    OutputFrequencyNominal  *prometheus.Desc
    OutputVoltage           *prometheus.Desc
    OutputVoltageNominal    *prometheus.Desc
    UPSBeeperStatus         *prometheus.Desc
    UPSDelayStart           *prometheus.Desc
    UPSDelayShutdown        *prometheus.Desc
    UPSStatus               *prometheus.Desc
    UPSType                 *prometheus.Desc

    ss StatusSource
}

var _ StatusSource = &GetUPSStatus{}
type StatusSource interface {
    GetUPSStatus()
}

func NewUPSCollector(ss StatusSource) *UPSCollector {
    labels := []string{"ups"}

    return &UPSCollector{
        Info: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","info"),
            "Metadata about a given UPS.",
            []string{"ups","hostname","model"},
            nil,
        ),

        BatteryChargePercent: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_battery_charge"),
            "Current UPS load percentage.",
            labels,
            nil,
        ),

        BatteryVoltage: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_battery_voltage"),
            "Current UPS battery voltage",
            labels,
            nil,
        ),

        BatteryVoltageLow: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_battery_voltage_low"),
            "Low UPS battery voltage",
            labels,
            nil,
        ),

        BatteryVoltageHigh: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_battery_voltage_high"),
            "High UPS battery voltage",
            labels,
            nil,
        ),
        BatteryVoltageNominal: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_battery_voltage_nominal"),
            "Nominal UPS battery voltage",
            labels,
            nil,
        ),

        InputVoltage: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_input_voltage"),
            "Current input line voltage",
            labels,
            nil,
        ),

        InputVoltageFault: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_input_voltage_fault"),
            "Fault input line voltage",
            labels,
            nil,
        ),


        OutputCurrentNominal: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_output_current_nominal"),
            "Output current nominal",
            labels,
            nil,
        ),

        OutputFrequency: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_output_frequency"),
            "Output frequency",
            labels,
            nil,
        ),

        OutputFrequencyNominal: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_output_frequency_nominal"),
            "Output nominal frequency",
            labels,
            nil,
        ),


        OutputVoltage: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_output_voltage"),
            "Output voltage",
            labels,
            nil,
        ),


        OutputVoltageNominal: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_output_voltage_nominal"),
            "Output nominal voltage",
            labels,
            nil,
        ),


        UPSBeeperStatus: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_ups_beeper_status"),
            "UPS beeper status",
            labels,
            nil,
        ),

        UPSDelayStart: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_ups_delay_start"),
            "UPS delay start",
            labels,
            nil,
        ),

        UPSDelayShutdown: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_ups_delay_shutdown"),
            "UPS delay shutdown",
            labels,
            nil,
        ),

        UPSStatus: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_ups_status"),
            "UPS Status",
            labels,
            nil,
        ),

        UPSType: prometheus.NewDesc(
            prometheus.BuildFQName(namespace,"","ups_ups_type"),
            "UPS status type",
            labels,
            nil,
        ),

        ss: ss,


    }
}



func (c *UPSCollector) Describe(ch chan<- *prometheus.Desc) {
    ds := []*prometheus.Desc{
        c.Info,
		c.BatteryChargePercent,
		c.BatteryVoltage,
		c.BatteryVoltageLow,
		c.BatteryVoltageHigh,
		c.BatteryVoltageNominal,
		c.InputVoltage,
		c.InputVoltageFault,
		c.OutputCurrentNominal,
		c.OutputFrequency,
		c.OutputFrequencyNominal,
		c.OutputVoltage,
		c.OutputVoltageNominal,
		c.UPSBeeperStatus,
		c.UPSDelayStart,
		c.UPSDelayShutdown,
		c.UPSStatus,
		c.UPSType,
    }

    for _, d := range ds{
        ch <- d
    }
}

func (c *UPSCollector) Collect(ch chan<- prometheus.Metric) {
    s, err := c.ss.GetUPSStatus()
    if err != nil {
        log.Printf("Failed collecting UPS metrics: %v", err)
        ch <- prometheus.NewInvalidMetric(c.Info,err)
        return
    }

    ch <- prometheus.MustNewConstMetric(
        c.Info,
        prometheus.GaugeValue,
        1,
        *nutUpsName, *nutAddr, s.ups_driver_version,
    )

    ch <- prometheus.MustNewConstMetric(
		c.BatteryChargePercent,
        prometheus.GaugeValue,
		s.ups_battery_charge,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.BatteryVoltage,
        prometheus.GaugeValue,
        s.ups_battery_voltage,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
        c.BatteryVoltageLow,
        prometheus.GaugeValue,
        s.ups_battery_voltage_low,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.BatteryVoltageHigh,
        prometheus.GaugeValue,
        s.ups_battery_voltage_high,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.BatteryVoltageNominal,
        prometheus.GaugeValue,
        s.ups_battery_voltage_nominal,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.InputVoltage,
        prometheus.GaugeValue,
        s.ups_input_voltage,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.InputVoltageFault,
        prometheus.GaugeValue,
        s.ups_input_voltage_fault,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.OutputCurrentNominal,
        prometheus.GaugeValue,
        s.ups_output_current_nominal,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.OutputFrequency,
        prometheus.GaugeValue,
        s.ups_output_frequency,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.OutputFrequencyNominal,
        prometheus.GaugeValue,
        s.ups_output_frequency_nominal,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.OutputVoltage,
        prometheus.GaugeValue,
        s.ups_output_voltage,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.OutputVoltageNominal,
        prometheus.GaugeValue,
        s.ups_output_voltage_nominal,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.UPSBeeperStatus,
        prometheus.GaugeValue,
        s.ups_ups_beeper_status,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.UPSDelayStart,
        prometheus.GaugeValue,
        s.ups_ups_delay_start,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.UPSDelayShutdown,
        prometheus.GaugeValue,
        s.ups_ups_delay_shutdown,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.UPSStatus,
        prometheus.GaugeValue,
        s.ups_ups_status,
        *nutUpsName,
    )

    ch <- prometheus.MustNewConstMetric(
		c.UPSType,
        prometheus.GaugeValue,
        s.ups_ups_type,
        *nutUpsName,
    )

}




