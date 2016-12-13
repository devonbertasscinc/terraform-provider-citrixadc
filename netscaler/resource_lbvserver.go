package netscaler

import (
	"github.com/chiradeep/go-nitro/config/lb"
	"github.com/chiradeep/go-nitro/config/ssl"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerLbvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbvserverFunc,
		Read:          readLbvserverFunc,
		Update:        updateLbvserverFunc,
		Delete:        deleteLbvserverFunc,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authn401": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnvsname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backuppersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"backupvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bypassaaaa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connfailover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datalength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dataoffset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dbslb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disableprimaryondown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healththreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpvsrresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertvserveripport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2conn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpolicy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpriority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"m": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macmoderetainvlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxautoscalemembers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minautoscalemembers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mssqlserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mysqlcharacterset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlprotocolversion": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlservercapabilities": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newservicerequest": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"newservicerequestincrementinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"newservicerequestunit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistencebackup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistencetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushlabel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushmulticlients": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"recursionavailable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectportrewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirurlflags": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"resrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtspnat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionless": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skippersistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sobackupaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"somethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sothreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v6netmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v6persistmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslcertkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In createLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var lbvserverName string
	if v, ok := d.GetOk("name"); ok {
		lbvserverName = v.(string)
	} else {
		lbvserverName = resource.PrefixedUniqueId("tf-lbvserver-")
		d.Set("name", lbvserverName)
	}
	sslcertkey, sok := d.GetOk("sslcertkey")
	if sok {
		exists := client.ResourceExists(netscaler.Sslcertkey.Type(), sslcertkey.(string))
		if !exists {
			return fmt.Errorf("[ERROR] netscaler-provider: Specified ssl cert key does not exist on netscaler!")
		}
	}

	lbvserver := lb.Lbvserver{
		Name:                               lbvserverName,
		Appflowlog:                         d.Get("appflowlog").(string),
		Authentication:                     d.Get("authentication").(string),
		Authenticationhost:                 d.Get("authenticationhost").(string),
		Authn401:                           d.Get("authn401").(string),
		Authnprofile:                       d.Get("authnprofile").(string),
		Authnvsname:                        d.Get("authnvsname").(string),
		Backuppersistencetimeout:           d.Get("backuppersistencetimeout").(int),
		Backupvserver:                      d.Get("backupvserver").(string),
		Bypassaaaa:                         d.Get("bypassaaaa").(string),
		Cacheable:                          d.Get("cacheable").(string),
		Clttimeout:                         d.Get("clttimeout").(int),
		Comment:                            d.Get("comment").(string),
		Connfailover:                       d.Get("connfailover").(string),
		Cookiename:                         d.Get("cookiename").(string),
		Datalength:                         d.Get("datalength").(int),
		Dataoffset:                         d.Get("dataoffset").(int),
		Dbprofilename:                      d.Get("dbprofilename").(string),
		Dbslb:                              d.Get("dbslb").(string),
		Disableprimaryondown:               d.Get("disableprimaryondown").(string),
		Dns64:                              d.Get("dns64").(string),
		Downstateflush:                     d.Get("downstateflush").(string),
		Hashlength:                         d.Get("hashlength").(int),
		Healththreshold:                    d.Get("healththreshold").(int),
		Httpprofilename:                    d.Get("httpprofilename").(string),
		Icmpvsrresponse:                    d.Get("icmpvsrresponse").(string),
		Insertvserveripport:                d.Get("insertvserveripport").(string),
		Ipmask:                             d.Get("ipmask").(string),
		Ippattern:                          d.Get("ippattern").(string),
		Ipv46:                              d.Get("ipv46").(string),
		L2conn:                             d.Get("l2conn").(string),
		Lbmethod:                           d.Get("lbmethod").(string),
		Listenpolicy:                       d.Get("listenpolicy").(string),
		Listenpriority:                     d.Get("listenpriority").(int),
		M:                                  d.Get("m").(string),
		Macmoderetainvlan:                  d.Get("macmoderetainvlan").(string),
		Maxautoscalemembers:                d.Get("maxautoscalemembers").(int),
		Minautoscalemembers:                d.Get("minautoscalemembers").(int),
		Mssqlserverversion:                 d.Get("mssqlserverversion").(string),
		Mysqlcharacterset:                  d.Get("mysqlcharacterset").(int),
		Mysqlprotocolversion:               d.Get("mysqlprotocolversion").(int),
		Mysqlservercapabilities:            d.Get("mysqlservercapabilities").(int),
		Mysqlserverversion:                 d.Get("mysqlserverversion").(string),
		Netmask:                            d.Get("netmask").(string),
		Netprofile:                         d.Get("netprofile").(string),
		Newname:                            d.Get("newname").(string),
		Newservicerequest:                  d.Get("newservicerequest").(int),
		Newservicerequestincrementinterval: d.Get("newservicerequestincrementinterval").(int),
		Newservicerequestunit:              d.Get("newservicerequestunit").(string),
		Persistencebackup:                  d.Get("persistencebackup").(string),
		Persistencetype:                    d.Get("persistencetype").(string),
		Persistmask:                        d.Get("persistmask").(string),
		Port:                               d.Get("port").(int),
		Pq:                                 d.Get("pq").(string),
		Push:                               d.Get("push").(string),
		Pushlabel:                          d.Get("pushlabel").(string),
		Pushmulticlients:                   d.Get("pushmulticlients").(string),
		Pushvserver:                        d.Get("pushvserver").(string),
		Range:                              d.Get("range").(int),
		Recursionavailable:                 d.Get("recursionavailable").(string),
		Redirectportrewrite:                d.Get("redirectportrewrite").(string),
		Redirurl:                           d.Get("redirurl").(string),
		Redirurlflags:                      d.Get("redirurlflags").(bool),
		Resrule:                            d.Get("resrule").(string),
		Rtspnat:                            d.Get("rtspnat").(string),
		Rule:                               d.Get("rule").(string),
		Sc:                                 d.Get("sc").(string),
		Servicename:                        d.Get("servicename").(string),
		Servicetype:                        d.Get("servicetype").(string),
		Sessionless:                        d.Get("sessionless").(string),
		Skippersistency:                    d.Get("skippersistency").(string),
		Sobackupaction:                     d.Get("sobackupaction").(string),
		Somethod:                           d.Get("somethod").(string),
		Sopersistence:                      d.Get("sopersistence").(string),
		Sopersistencetimeout:               d.Get("sopersistencetimeout").(int),
		Sothreshold:                        d.Get("sothreshold").(int),
		State:                              d.Get("state").(string),
		Tcpprofilename:                     d.Get("tcpprofilename").(string),
		Td:                                 d.Get("td").(int),
		Timeout:                            d.Get("timeout").(int),
		Tosid:                              d.Get("tosid").(int),
		V6netmasklen:                       d.Get("v6netmasklen").(int),
		V6persistmasklen:                   d.Get("v6persistmasklen").(int),
		Vipheader:                          d.Get("vipheader").(string),
		Weight:                             d.Get("weight").(int),
	}

	_, err := client.AddResource(netscaler.Lbvserver.Type(), lbvserverName, &lbvserver)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: could not add resource %s of type %s", netscaler.Lbvserver.Type(), lbvserverName)
		return err
	}
	if sok { //ssl cert is specified
		binding := ssl.Sslvserversslcertkeybinding{
			Vservername: lbvserverName,
			Certkeyname: sslcertkey.(string),
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl cert %s to lbvserver %s", sslcertkey, lbvserverName)
		err = client.BindResource(netscaler.Sslvserver.Type(), lbvserverName, netscaler.Sslcertkey.Type(), sslcertkey.(string), &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkey, lbvserverName)
			err2 := client.DeleteResource(netscaler.Lbvserver.Type(), lbvserverName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete lbvserver %s after bind to ssl cert failed", lbvserverName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to delete lbvserver %s after bind to ssl cert failed", lbvserverName)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkey, lbvserverName)
		}
	}

	d.SetId(lbvserverName)

	err = readLbvserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbvserver but we can't read it ?? %s", lbvserverName)
		return nil
	}
	return nil
}

func readLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading lbvserver state %s", lbvserverName)
	data, err := client.FindResource(netscaler.Lbvserver.Type(), lbvserverName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing lbvserver state %s", lbvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("authentication", data["authentication"])
	d.Set("authenticationhost", data["authenticationhost"])
	d.Set("authn401", data["authn401"])
	d.Set("authnprofile", data["authnprofile"])
	d.Set("authnvsname", data["authnvsname"])
	d.Set("backuppersistencetimeout", data["backuppersistencetimeout"])
	d.Set("backupvserver", data["backupvserver"])
	d.Set("bypassaaaa", data["bypassaaaa"])
	d.Set("cacheable", data["cacheable"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("comment", data["comment"])
	d.Set("connfailover", data["connfailover"])
	d.Set("cookiename", data["cookiename"])
	d.Set("datalength", data["datalength"])
	d.Set("dataoffset", data["dataoffset"])
	d.Set("dbprofilename", data["dbprofilename"])
	d.Set("dbslb", data["dbslb"])
	d.Set("disableprimaryondown", data["disableprimaryondown"])
	d.Set("dns64", data["dns64"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("hashlength", data["hashlength"])
	d.Set("healththreshold", data["healththreshold"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("icmpvsrresponse", data["icmpvsrresponse"])
	d.Set("insertvserveripport", data["insertvserveripport"])
	d.Set("ipmask", data["ipmask"])
	d.Set("ippattern", data["ippattern"])
	d.Set("ipv46", data["ipv46"])
	d.Set("l2conn", data["l2conn"])
	d.Set("lbmethod", data["lbmethod"])
	d.Set("listenpolicy", data["listenpolicy"])
	d.Set("listenpriority", data["listenpriority"])
	d.Set("m", data["m"])
	d.Set("macmoderetainvlan", data["macmoderetainvlan"])
	d.Set("maxautoscalemembers", data["maxautoscalemembers"])
	d.Set("minautoscalemembers", data["minautoscalemembers"])
	d.Set("mssqlserverversion", data["mssqlserverversion"])
	d.Set("mysqlcharacterset", data["mysqlcharacterset"])
	d.Set("mysqlprotocolversion", data["mysqlprotocolversion"])
	d.Set("mysqlservercapabilities", data["mysqlservercapabilities"])
	d.Set("mysqlserverversion", data["mysqlserverversion"])
	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])
	d.Set("netprofile", data["netprofile"])
	d.Set("newname", data["newname"])
	d.Set("newservicerequest", data["newservicerequest"])
	d.Set("newservicerequestincrementinterval", data["newservicerequestincrementinterval"])
	d.Set("newservicerequestunit", data["newservicerequestunit"])
	d.Set("persistencebackup", data["persistencebackup"])
	d.Set("persistencetype", data["persistencetype"])
	d.Set("persistmask", data["persistmask"])
	d.Set("port", data["port"])
	d.Set("pq", data["pq"])
	d.Set("push", data["push"])
	d.Set("pushlabel", data["pushlabel"])
	d.Set("pushmulticlients", data["pushmulticlients"])
	d.Set("pushvserver", data["pushvserver"])
	d.Set("range", data["range"])
	d.Set("recursionavailable", data["recursionavailable"])
	d.Set("redirectportrewrite", data["redirectportrewrite"])
	d.Set("redirurl", data["redirurl"])
	d.Set("redirurlflags", data["redirurlflags"])
	d.Set("resrule", data["resrule"])
	d.Set("rtspnat", data["rtspnat"])
	d.Set("rule", data["rule"])
	d.Set("sc", data["sc"])
	d.Set("servicename", data["servicename"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sessionless", data["sessionless"])
	d.Set("skippersistency", data["skippersistency"])
	d.Set("sobackupaction", data["sobackupaction"])
	d.Set("somethod", data["somethod"])
	d.Set("sopersistence", data["sopersistence"])
	d.Set("sopersistencetimeout", data["sopersistencetimeout"])
	d.Set("sothreshold", data["sothreshold"])
	d.Set("state", data["state"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("timeout", data["timeout"])
	d.Set("tosid", data["tosid"])
	d.Set("v6netmasklen", data["v6netmasklen"])
	d.Set("v6persistmasklen", data["v6persistmasklen"])
	d.Set("vipheader", data["vipheader"])
	d.Set("weight", data["weight"])

	bindings, err := client.FindAllBoundResources(netscaler.Sslvserver.Type(), lbvserverName, netscaler.Sslcertkey.Type())
	if err != nil {
		log.Printf("[WARN] netscaler-provider: lbvserver binding to ssl error %s", lbvserverName)
		return nil
	}
	var boundCert string
	for _, binding := range bindings {
		cert, ok := binding["certkeyname"]
		if ok {
			boundCert = cert.(string)
			break
		}
	}
	d.Set("sslcertkey", boundCert)

	return nil

}

func updateLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In updateLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Get("name").(string)

	lbvserver := lb.Lbvserver{
		Name: d.Get("name").(string),
	}
	hasChange := false
	sslcertkeyChanged := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG] netscaler-provider:  Appflowlog has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG] netscaler-provider:  Authentication has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authenticationhost") {
		log.Printf("[DEBUG] netscaler-provider:  Authenticationhost has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authenticationhost = d.Get("authenticationhost").(string)
		hasChange = true
	}
	if d.HasChange("authn401") {
		log.Printf("[DEBUG] netscaler-provider:  Authn401 has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authn401 = d.Get("authn401").(string)
		hasChange = true
	}
	if d.HasChange("authnprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Authnprofile has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authnprofile = d.Get("authnprofile").(string)
		hasChange = true
	}
	if d.HasChange("authnvsname") {
		log.Printf("[DEBUG] netscaler-provider:  Authnvsname has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Authnvsname = d.Get("authnvsname").(string)
		hasChange = true
	}
	if d.HasChange("backuppersistencetimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Backuppersistencetimeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Backuppersistencetimeout = d.Get("backuppersistencetimeout").(int)
		hasChange = true
	}
	if d.HasChange("backupvserver") {
		log.Printf("[DEBUG] netscaler-provider:  Backupvserver has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Backupvserver = d.Get("backupvserver").(string)
		hasChange = true
	}
	if d.HasChange("bypassaaaa") {
		log.Printf("[DEBUG] netscaler-provider:  Bypassaaaa has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Bypassaaaa = d.Get("bypassaaaa").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG] netscaler-provider:  Cacheable has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Clttimeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG] netscaler-provider:  Comment has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("connfailover") {
		log.Printf("[DEBUG] netscaler-provider:  Connfailover has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Connfailover = d.Get("connfailover").(string)
		hasChange = true
	}
	if d.HasChange("cookiename") {
		log.Printf("[DEBUG] netscaler-provider:  Cookiename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Cookiename = d.Get("cookiename").(string)
		hasChange = true
	}
	if d.HasChange("datalength") {
		log.Printf("[DEBUG] netscaler-provider:  Datalength has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Datalength = d.Get("datalength").(int)
		hasChange = true
	}
	if d.HasChange("dataoffset") {
		log.Printf("[DEBUG] netscaler-provider:  Dataoffset has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dataoffset = d.Get("dataoffset").(int)
		hasChange = true
	}
	if d.HasChange("dbprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Dbprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dbprofilename = d.Get("dbprofilename").(string)
		hasChange = true
	}
	if d.HasChange("dbslb") {
		log.Printf("[DEBUG] netscaler-provider:  Dbslb has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dbslb = d.Get("dbslb").(string)
		hasChange = true
	}
	if d.HasChange("disableprimaryondown") {
		log.Printf("[DEBUG] netscaler-provider:  Disableprimaryondown has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
		hasChange = true
	}
	if d.HasChange("dns64") {
		log.Printf("[DEBUG] netscaler-provider:  Dns64 has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Dns64 = d.Get("dns64").(string)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG] netscaler-provider:  Downstateflush has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("hashlength") {
		log.Printf("[DEBUG] netscaler-provider:  Hashlength has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Hashlength = d.Get("hashlength").(int)
		hasChange = true
	}
	if d.HasChange("healththreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Healththreshold has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Healththreshold = d.Get("healththreshold").(int)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Httpprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("icmpvsrresponse") {
		log.Printf("[DEBUG] netscaler-provider:  Icmpvsrresponse has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
		hasChange = true
	}
	if d.HasChange("insertvserveripport") {
		log.Printf("[DEBUG] netscaler-provider:  Insertvserveripport has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Insertvserveripport = d.Get("insertvserveripport").(string)
		hasChange = true
	}
	if d.HasChange("ipmask") {
		log.Printf("[DEBUG] netscaler-provider:  Ipmask has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ipmask = d.Get("ipmask").(string)
		hasChange = true
	}
	if d.HasChange("ippattern") {
		log.Printf("[DEBUG] netscaler-provider:  Ippattern has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ippattern = d.Get("ippattern").(string)
		hasChange = true
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG] netscaler-provider:  Ipv46 has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Ipv46 = d.Get("ipv46").(string)
		hasChange = true
	}
	if d.HasChange("l2conn") {
		log.Printf("[DEBUG] netscaler-provider:  L2conn has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.L2conn = d.Get("l2conn").(string)
		hasChange = true
	}
	if d.HasChange("lbmethod") {
		log.Printf("[DEBUG] netscaler-provider:  Lbmethod has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Lbmethod = d.Get("lbmethod").(string)
		hasChange = true
	}
	if d.HasChange("listenpolicy") {
		log.Printf("[DEBUG] netscaler-provider:  Listenpolicy has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Listenpolicy = d.Get("listenpolicy").(string)
		hasChange = true
	}
	if d.HasChange("listenpriority") {
		log.Printf("[DEBUG] netscaler-provider:  Listenpriority has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Listenpriority = d.Get("listenpriority").(int)
		hasChange = true
	}
	if d.HasChange("m") {
		log.Printf("[DEBUG] netscaler-provider:  M has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.M = d.Get("m").(string)
		hasChange = true
	}
	if d.HasChange("macmoderetainvlan") {
		log.Printf("[DEBUG] netscaler-provider:  Macmoderetainvlan has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Macmoderetainvlan = d.Get("macmoderetainvlan").(string)
		hasChange = true
	}
	if d.HasChange("maxautoscalemembers") {
		log.Printf("[DEBUG] netscaler-provider:  Maxautoscalemembers has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Maxautoscalemembers = d.Get("maxautoscalemembers").(int)
		hasChange = true
	}
	if d.HasChange("minautoscalemembers") {
		log.Printf("[DEBUG] netscaler-provider:  Minautoscalemembers has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Minautoscalemembers = d.Get("minautoscalemembers").(int)
		hasChange = true
	}
	if d.HasChange("mssqlserverversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mssqlserverversion has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mssqlserverversion = d.Get("mssqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("mysqlcharacterset") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlcharacterset has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlcharacterset = d.Get("mysqlcharacterset").(int)
		hasChange = true
	}
	if d.HasChange("mysqlprotocolversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlprotocolversion has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlprotocolversion = d.Get("mysqlprotocolversion").(int)
		hasChange = true
	}
	if d.HasChange("mysqlservercapabilities") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlservercapabilities has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlservercapabilities = d.Get("mysqlservercapabilities").(int)
		hasChange = true
	}
	if d.HasChange("mysqlserverversion") {
		log.Printf("[DEBUG] netscaler-provider:  Mysqlserverversion has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Mysqlserverversion = d.Get("mysqlserverversion").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG] netscaler-provider:  Name has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG] netscaler-provider:  Netmask has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Netmask = d.Get("netmask").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Netprofile has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG] netscaler-provider:  Newname has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("newservicerequest") {
		log.Printf("[DEBUG] netscaler-provider:  Newservicerequest has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Newservicerequest = d.Get("newservicerequest").(int)
		hasChange = true
	}
	if d.HasChange("newservicerequestincrementinterval") {
		log.Printf("[DEBUG] netscaler-provider:  Newservicerequestincrementinterval has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Newservicerequestincrementinterval = d.Get("newservicerequestincrementinterval").(int)
		hasChange = true
	}
	if d.HasChange("newservicerequestunit") {
		log.Printf("[DEBUG] netscaler-provider:  Newservicerequestunit has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Newservicerequestunit = d.Get("newservicerequestunit").(string)
		hasChange = true
	}
	if d.HasChange("persistencebackup") {
		log.Printf("[DEBUG] netscaler-provider:  Persistencebackup has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Persistencebackup = d.Get("persistencebackup").(string)
		hasChange = true
	}
	if d.HasChange("persistencetype") {
		log.Printf("[DEBUG] netscaler-provider:  Persistencetype has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Persistencetype = d.Get("persistencetype").(string)
		hasChange = true
	}
	if d.HasChange("persistmask") {
		log.Printf("[DEBUG] netscaler-provider:  Persistmask has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Persistmask = d.Get("persistmask").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG] netscaler-provider:  Port has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("pq") {
		log.Printf("[DEBUG] netscaler-provider:  Pq has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Pq = d.Get("pq").(string)
		hasChange = true
	}
	if d.HasChange("push") {
		log.Printf("[DEBUG] netscaler-provider:  Push has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Push = d.Get("push").(string)
		hasChange = true
	}
	if d.HasChange("pushlabel") {
		log.Printf("[DEBUG] netscaler-provider:  Pushlabel has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Pushlabel = d.Get("pushlabel").(string)
		hasChange = true
	}
	if d.HasChange("pushmulticlients") {
		log.Printf("[DEBUG] netscaler-provider:  Pushmulticlients has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Pushmulticlients = d.Get("pushmulticlients").(string)
		hasChange = true
	}
	if d.HasChange("pushvserver") {
		log.Printf("[DEBUG] netscaler-provider:  Pushvserver has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Pushvserver = d.Get("pushvserver").(string)
		hasChange = true
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG] netscaler-provider:  Range has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Range = d.Get("range").(int)
		hasChange = true
	}
	if d.HasChange("recursionavailable") {
		log.Printf("[DEBUG] netscaler-provider:  Recursionavailable has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Recursionavailable = d.Get("recursionavailable").(string)
		hasChange = true
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG] netscaler-provider:  Redirectportrewrite has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Redirectportrewrite = d.Get("redirectportrewrite").(string)
		hasChange = true
	}
	if d.HasChange("redirurl") {
		log.Printf("[DEBUG] netscaler-provider:  Redirurl has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Redirurl = d.Get("redirurl").(string)
		hasChange = true
	}
	if d.HasChange("redirurlflags") {
		log.Printf("[DEBUG] netscaler-provider:  Redirurlflags has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Redirurlflags = d.Get("redirurlflags").(bool)
		hasChange = true
	}
	if d.HasChange("resrule") {
		log.Printf("[DEBUG] netscaler-provider:  Resrule has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Resrule = d.Get("resrule").(string)
		hasChange = true
	}
	if d.HasChange("rtspnat") {
		log.Printf("[DEBUG] netscaler-provider:  Rtspnat has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Rtspnat = d.Get("rtspnat").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG] netscaler-provider:  Rule has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("sc") {
		log.Printf("[DEBUG] netscaler-provider:  Sc has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sc = d.Get("sc").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG] netscaler-provider:  Servicename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG] netscaler-provider:  Servicetype has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sessionless") {
		log.Printf("[DEBUG] netscaler-provider:  Sessionless has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sessionless = d.Get("sessionless").(string)
		hasChange = true
	}
	if d.HasChange("skippersistency") {
		log.Printf("[DEBUG] netscaler-provider:  Skippersistency has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Skippersistency = d.Get("skippersistency").(string)
		hasChange = true
	}
	if d.HasChange("sobackupaction") {
		log.Printf("[DEBUG] netscaler-provider:  Sobackupaction has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sobackupaction = d.Get("sobackupaction").(string)
		hasChange = true
	}
	if d.HasChange("somethod") {
		log.Printf("[DEBUG] netscaler-provider:  Somethod has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Somethod = d.Get("somethod").(string)
		hasChange = true
	}
	if d.HasChange("sopersistence") {
		log.Printf("[DEBUG] netscaler-provider:  Sopersistence has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sopersistence = d.Get("sopersistence").(string)
		hasChange = true
	}
	if d.HasChange("sopersistencetimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Sopersistencetimeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sopersistencetimeout = d.Get("sopersistencetimeout").(int)
		hasChange = true
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Sothreshold has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Sothreshold = d.Get("sothreshold").(int)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG] netscaler-provider:  State has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Tcpprofilename has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG] netscaler-provider:  Td has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG] netscaler-provider:  Timeout has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("tosid") {
		log.Printf("[DEBUG] netscaler-provider:  Tosid has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Tosid = d.Get("tosid").(int)
		hasChange = true
	}
	if d.HasChange("v6netmasklen") {
		log.Printf("[DEBUG] netscaler-provider:  V6netmasklen has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.V6netmasklen = d.Get("v6netmasklen").(int)
		hasChange = true
	}
	if d.HasChange("v6persistmasklen") {
		log.Printf("[DEBUG] netscaler-provider:  V6persistmasklen has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.V6persistmasklen = d.Get("v6persistmasklen").(int)
		hasChange = true
	}
	if d.HasChange("vipheader") {
		log.Printf("[DEBUG] netscaler-provider:  Vipheader has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Vipheader = d.Get("vipheader").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG] netscaler-provider:  Weight has changed for lbvserver %s, starting update", lbvserverName)
		lbvserver.Weight = d.Get("weight").(int)
		hasChange = true
	}
	if d.HasChange("sslcertkey") {
		log.Printf("[DEBUG] netscaler-provider:  ssl certkey has changed for lbvserver %s, starting update", lbvserverName)
		sslcertkeyChanged = true
	}

	sslcertkey := d.Get("sslcertkey")
	sslcertkeyName := sslcertkey.(string)
	if sslcertkeyChanged {
		//Binding has to be updated
		//First we unbind from lb vserver
		oldSslcertkey, _ := d.GetChange("sslcertkey")
		oldSslcertkeyName := oldSslcertkey.(string)
		if oldSslcertkeyName != "" {
			err := client.UnbindResource(netscaler.Sslvserver.Type(), lbvserverName, netscaler.Sslcertkey.Type(), oldSslcertkeyName, "certkeyname")
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding sslcertkey from lbvserver %s", oldSslcertkeyName)
			}
			log.Printf("[DEBUG] netscaler-provider: sslcertkey has been unbound from lbvserver for sslcertkey %s ", oldSslcertkeyName)
		}
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Lbvserver.Type(), lbvserverName, &lbvserver)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error updating lbvserver %s", lbvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: lbvserver has been updated  lbvserver %s ", lbvserverName)
	}

	if sslcertkeyChanged && sslcertkeyName != "" {
		//Binding has to be updated
		//rebind
		binding := ssl.Sslvserversslcertkeybinding{
			Vservername: lbvserverName,
			Certkeyname: sslcertkeyName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding ssl cert %s to lbvserver %s", sslcertkeyName, lbvserverName)
		err := client.BindResource(netscaler.Sslvserver.Type(), lbvserverName, netscaler.Sslcertkey.Type(), sslcertkeyName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkeyName, lbvserverName)
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind ssl cert %s to lbvserver %s", sslcertkeyName, lbvserverName)
		}
		log.Printf("[DEBUG] netscaler-provider: new ssl cert has been bound to lbvserver  sslcertkey %s lbvserver %s", sslcertkeyName, lbvserverName)
	}
	return readLbvserverFunc(d, meta)
}

func deleteLbvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In deleteLbvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Id()
	err := client.DeleteResource(netscaler.Lbvserver.Type(), lbvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}