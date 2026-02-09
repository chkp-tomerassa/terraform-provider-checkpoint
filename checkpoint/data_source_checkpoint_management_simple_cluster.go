package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementSimpleCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSimpleClusterRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"cluster_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster mode.",
			},
			"geo_mode": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Cluster High Availability Geo mode. This setting applies only to a cluster deployed in a cloud. Available when the cluster mode equals \"cluster-xl-ha\".",
			},
			"advanced_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_persistence": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Handling established connections when installing a new policy.",
						},
						"sam": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "SAM.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"forward_to_other_sam_servers": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Forward SAM clients' requests to other SAM servers.",
									},
									"use_early_versions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Use early versions compatibility mode.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Use early versions compatibility mode.",
												},
												"compatibility_mode": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Early versions compatibility mode.",
												},
											},
										},
									},
									"purge_sam_file": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Purge SAM File.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Purge SAM File.",
												},
												"purge_when_size_reaches_to": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Purge SAM File When it Reaches to.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"enable_https_inspection": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.",
			},
			"fetch_policy": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Security management server(s) to fetch the policy from.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hit_count": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hit count tracks the number of connections each rule matches.",
			},
			"https_inspection": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "HTTPS inspection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_on_failure": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to bypass all requests (Fail-open) in case of internal system error.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"site_categorization_allow_mode": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to 'background' in order to allowed requests until categorization is complete.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_untrusted_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with untrusted server certificate.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_revoked_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_expired_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with expired server certificate.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
					},
				},
			},
			"identity_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Identity awareness blade enabled.",
			},
			"identity_awareness_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gateway Identity Awareness settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"browser_based_authentication": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Browser Based Authentication source.",
						},
						"browser_based_authentication_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Browser Based Authentication settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication Settings for Browser Based Authentication.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Authentication method.",
												},
												"identity_provider": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Identity provider object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"identity provider\".",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"radius": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Radius server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"radius\".",
												},
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "External user profile.",
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Internal users.",
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Users from external directories.",
															},
															"specific": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "LDAP AU objects identified by the name or UID. Must be set when \"users-from-external-directories\" was selected to be \"specific\".",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"browser_based_authentication_portal_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Browser Based Authentication portal settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"portal_web_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal web settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"aliases": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "List of URL aliases that are redirected to the main portal URL.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"main_url": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The main URL for the web portal.",
															},
														},
													},
												},
												"certificate_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal certificate settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"base64_certificate": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
															},
															"base64_password": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Password (encoded in Base64 with padding) for the certificate file.",
															},
														},
													},
												},
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"identity_agent": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Identity Agent source.",
						},
						"identity_agent_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity Agent settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agents_interval_keepalive": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Agents send keepalive period (minutes).",
									},
									"user_reauthenticate_interval": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Agent reauthenticate time interval (minutes).",
									},
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication Settings for Identity Agent.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Authentication method.",
												},
												"radius": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Radius server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"radius\".",
												},
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "External user profile.",
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Internal users.",
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Users from external directories.",
															},
															"specific": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "LDAP AU objects identified by the name or UID. Must be set when \"users-from-external-directories\" was selected to be \"specific\".",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"identity_agent_portal_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Identity Agent accessibility settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"identity_collector": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Identity Collector source.",
						},
						"identity_collector_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity Collector settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authorized_clients": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authorized Clients.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"client": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Host / Network Group Name or UID.",
												},
												"client_secret": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Client Secret.",
												},
											},
										},
									},
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication Settings for Identity Collector.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "External user profile.",
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Internal users.",
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Users from external directories.",
															},
															"specific": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "LDAP AU objects identified by the name or UID. Must be set when \"users-from-external-directories\" was selected to be \"specific\".",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"client_access_permissions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Identity Collector accessibility settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"identity_sharing_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity sharing settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"share_with_other_gateways": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable identity sharing with other gateways.",
									},
									"receive_from_other_gateways": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable receiving identity from other gateways.",
									},
									"receive_from": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Gateway(s) to receive identity from.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"proxy_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity-Awareness Proxy settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"detect_using_x_forward_for": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP.",
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Remote Access Identity source.",
						},
					},
				},
			},
			"ips_update_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.",
			},
			"nat_hide_internal_interfaces": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hide internal networks behind the Gateway's external IP.",
			},
			"nat_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to add automatic address translation rules.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hide behind method. This parameter is forbidden in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"platform_portal_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Platform portal settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"portal_web_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal web settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aliases": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of URL aliases that are redirected to the main portal URL.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"main_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The main URL for the web portal.",
									},
								},
							},
						},
						"certificate_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal certificate settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base64_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
									},
									"base64_password": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Password (encoded in Base64 with padding) for the certificate file.",
									},
								},
							},
						},
						"accessibility": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal access settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_access_from": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allowed access to the web portal (based on interfaces, or security policy).",
									},
									"internal_access_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Configuration of the additional portal access settings for internal interfaces only.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"undefined": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
												},
												"dmz": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
												},
												"vpn": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"proxy_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Proxy Server for Gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_custom_proxy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use custom proxy settings for this network object.",
							//Default:     false,
						},
						"proxy_server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "N/A",
							//Default:     80,
						},
					},
				},
			},
			"qos": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "QoS.",
			},
			"usercheck_portal_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "UserCheck portal settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.",
						},
						"portal_web_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal web settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aliases": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of URL aliases that are redirected to the main portal URL.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"main_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The main URL for the web portal.",
									},
								},
							},
						},
						"certificate_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal certificate settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base64_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
									},
									"base64_password": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Password (encoded in Base64 with padding) for the certificate file.",
									},
								},
							},
						},
						"accessibility": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal access settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_access_from": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allowed access to the web portal (based on interfaces, or security policy).",
									},
									"internal_access_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Configuration of the additional portal access settings for internal interfaces only.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"undefined": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
												},
												"dmz": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
												},
												"vpn": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"zero_phishing": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Zero Phishing blade enabled.",
			},
			"zero_phishing_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zero Phishing gateway FQDN.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"interface_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster interface type.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network address.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network mask length.",
						},
						"anti_spoofing": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Anti spoofing.",
						},
						"anti_spoofing_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Anti spoofing settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
									},
								},
							},
						},
						"multicast_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Multicast IP Address.",
						},
						"multicast_address_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Multicast Address Type.",
						},
						"security_zone": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Security zone.",
						},
						"security_zone_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Security zone settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_calculated": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Security Zone is calculated according to where the interface leads to.",
									},
									"specific_zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Security Zone specified manually.",
									},
								},
							},
						},
						"topology": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topology.",
						},
						"topology_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Topology settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether this interface leads to demilitarized zone (perimeter network).",
									},
									"ip_address_behind_this_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Ip address behind this interface.",
									},
									"specific_network": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network behind this interface.",
									},
								},
							},
						},
						"topology_automatic_calculation": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shows the automatic topology calculation.",
						},
						"color": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Color of the object. Should be one of existing colors.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comments string.",
						},
					},
				},
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cluster members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 or IPv6 address.",
						},
						"sic_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secure Internal Communication name.",
						},
						"sic_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secure Internal Communication state.",
						},
						"interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Network interfaces.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name. Should be unique in the domain.",
									},
									"ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 address.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv6 address.",
									},
									"ipv4_network_mask": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 network address.",
									},
									"ipv6_network_mask": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv6 network address.",
									},
									"ipv4_mask_length": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 network mask length.",
									},
									"ipv6_mask_length": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv6 network mask length.",
									},
								},
							},
						},
					},
				},
			},
			"anti_bot": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Anti-Bot blade enabled.",
			},
			"anti_virus": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Anti-Virus blade enabled.",
			},
			"application_control": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Application Control blade enabled.",
			},
			"content_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Content Awareness blade enabled.",
			},
			"data_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Data Awareness blade enabled.",
			},
			"firewall": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Firewall blade enabled.",
			},
			"firewall_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Firewall settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_calculate_connections_hash_table_size_and_memory_pool": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Auto calculate connections hash table size and memory pool.",
						},
						"auto_maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Auto maximum limit for concurrent connections.",
						},
						"connections_hash_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Connections hash size.",
						},
						"maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum limit for concurrent connections.",
						},
						"maximum_memory_pool_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum memory pool size.",
						},
						"memory_pool_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory pool size.",
						},
					},
				},
			},
			"ips": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Intrusion Prevention System blade enabled.",
			},
			"ips_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cluster IPS settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_all_under_load": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Disable/enable all IPS protections until CPU and memory levels are back to normal.",
						},
						"bypass_track_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Track options when all IPS protections are disabled until CPU/memory levels are back to normal.",
						},
						"top_cpu_consuming_protections": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Provides a way to reduce CPU levels on machines under load by disabling the top CPU consuming IPS protections.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disable_period": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Duration (in hours) for disabling the protections.",
									},
									"disable_under_load": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Temporarily disable/enable top CPU consuming IPS protections.",
									},
								},
							},
						},
						"activation_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether the IPS blade operates in Detect Only mode or enforces the configured IPS Policy.",
						},
						"cpu_usage_low_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "CPU usage low threshold percentage (1-99).",
						},
						"cpu_usage_high_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "CPU usage high threshold percentage (1-99).",
						},
						"memory_usage_low_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory usage low threshold percentage (1-99).",
						},
						"memory_usage_high_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory usage high threshold percentage (1-99).",
						},
						"send_threat_cloud_info": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Help improve Check Point Threat Prevention product by sending anonymous information.",
						},
						"reject_on_cluster_fail_over": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Define the IPS connections during fail over reject packets or accept packets.",
						},
					},
				},
			},
			"threat_emulation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Threat Emulation blade enabled.",
			},
			"url_filtering": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "URL Filtering blade enabled.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Dynamic IP address.",
			},
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS name.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster platform version.",
			},
			"hardware": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster platform hardware.",
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication name.",
			},
			"sic_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication state.",
			},
			"save_logs_locally": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Save logs locally.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Server(s) to send alerts to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Backup server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"logs_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Logs settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable alert when free disk space is below threshold.",
						},
						"alert_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below metrics.",
						},
						"alert_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Alert when free disk space below threshold.",
						},
						"alert_when_free_disk_space_below_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below type.",
						},
						"before_delete_keep_logs_from_the_last_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable before delete keep logs from the last days.",
						},
						"before_delete_keep_logs_from_the_last_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Before delete keep logs from the last days threshold.",
						},
						"before_delete_run_script": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Before delete run script.",
						},
						"before_delete_run_script_command": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Before delete run script command.",
						},
						"delete_index_files_older_than_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete index files older than days.",
						},
						"delete_index_files_older_than_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete index files older than days threshold.",
						},
						"delete_index_files_when_index_size_above": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete index files when index size above.",
						},
						"delete_index_files_when_index_size_above_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete index files when index size above metrics.",
						},
						"delete_index_files_when_index_size_above_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete index files when index size above threshold.",
						},
						"delete_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete when free disk space below.",
						},
						"delete_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete when free disk space below metric.",
						},
						"delete_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete when free disk space below threshold.",
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable detect new citrix ica application names.",
						},
						"forward_logs_to_log_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable forward logs to log server.",
						},
						"forward_logs_to_log_server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server name.",
						},
						"forward_logs_to_log_server_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server schedule name.",
						},
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Free disk space metrics.",
						},
						"perform_log_rotate_before_log_forwarding": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable perform log rotate before log forwarding.",
						},
						"reject_connections_when_free_disk_space_below_threshold": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable reject connections when free disk space below threshold.",
						},
						"reserve_for_packet_capture_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reserve for packet capture metrics.",
						},
						"reserve_for_packet_capture_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Reserve for packet capture threshold.",
						},
						"rotate_log_by_file_size": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log by file size.",
						},
						"rotate_log_file_size_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Log file size threshold.",
						},
						"rotate_log_on_schedule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log on schedule.",
						},
						"rotate_log_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rotate log schedule name.",
						},
						"stop_logging_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable stop logging when free disk space below.",
						},
						"stop_logging_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Stop logging when free disk space below threshold.",
						},
						"stop_logging_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Stop logging when free disk space below metrics.",
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable turn on qos logging.",
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Update account log in every amount of seconds.",
						},
					},
				},
			},
			"vpn": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "VPN blade enabled.",
			},
			"vpn_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gateway VPN settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Authentication.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_clients": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of VPN Authentication clients identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"link_selection": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Link Selection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_selection": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP selection",
									},
									"dns_resolving_hostname": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "DNS Resolving Hostname. Must be set when \"ip-selection\" was selected to be \"dns-resolving-from-hostname\".",
									},
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP Address. Must be set when \"ip-selection\" was selected to be \"use-selected-address-from-topology\" or \"use-statically-nated-ip\"",
									},
								},
							},
						},
						"maximum_concurrent_ike_negotiations": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum concurrent ike negotiations",
						},
						"maximum_concurrent_tunnels": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum concurrent tunnels",
						},
						"office_mode": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Office Mode Permissions. When selected to be \"off\", all the other definitions are irrelevant.",
									},
									"group": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Group. Identified by name or UID. Must be set when \"office-mode-permissions\" was selected to be \"group\".",
									},
									"allocate_ip_address_from": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_server": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Radius server used to authenticate the user.",
												},
												"use_allocate_method": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Use Allocate Method.",
												},
												"allocate_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Using either Manual (IP Pool) or Automatic (DHCP). Must be set when \"use-allocate-method\" is true.",
												},
												"manual_network": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Manual Network. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"manual\".",
												},
												"dhcp_server": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "DHCP Server. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"virtual_ip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Virtual IPV4 address for DHCP server replies. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"dhcp_mac_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Calculated MAC address for DHCP allocation. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"optional_parameters": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"use_primary_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Primary DNS Server.",
															},
															"primary_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Primary DNS Server. Identified by name or UID. Must be set when \"use-primary-dns-server\" is true and can not be set when \"use-primary-dns-server\" is false.",
															},
															"use_first_backup_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use First Backup DNS Server.",
															},
															"first_backup_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "First Backup DNS Server. Identified by name or UID. Must be set when \"use-first-backup-dns-server\" is true and can not be set when \"use-first-backup-dns-server\" is false.",
															},
															"use_second_backup_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Second Backup DNS Server.",
															},
															"second_backup_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Second Backup DNS Server. Identified by name or UID. Must be set when \"use-second-backup-dns-server\" is true and can not be set when \"use-second-backup-dns-server\" is false.",
															},
															"dns_suffixes": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "DNS Suffixes.",
															},
															"use_primary_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Primary WINS Server.",
															},
															"primary_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Primary WINS Server. Identified by name or UID. Must be set when \"use-primary-wins-server\" is true and can not be set when \"use-primary-wins-server\" is false.",
															},
															"use_first_backup_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use First Backup WINS Server.",
															},
															"first_backup_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "First Backup WINS Server. Identified by name or UID. Must be set when \"use-first-backup-wins-server\" is true and can not be set when \"use-first-backup-wins-server\" is false.",
															},
															"use_second_backup_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Second Backup WINS Server.",
															},
															"second_backup_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Second Backup WINS Server. Identified by name or UID. Must be set when \"use-second-backup-wins-server\" is true and can not be set when \"use-second-backup-wins-server\" is false.",
															},
															"ip_lease_duration": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "IP Lease Duration in Minutes. The value must be in the range 2-32767.",
															},
														},
													},
												},
											},
										},
									},
									"support_multiple_interfaces": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support connectivity enhancement for gateways with multiple external interfaces.",
									},
									"perform_anti_spoofing": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Perform Anti-Spoofing on Office Mode addresses.",
									},
									"anti_spoofing_additional_addresses": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when \"perform-anti-spoofings\" is true.",
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Remote Access.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"support_l2tp": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support L2TP (relevant only when office mode is active).",
									},
									"l2tp_auth_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "L2TP Authentication Method. Must be set when \"support-l2tp\" is true.",
									},
									"l2tp_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "L2TP Certificate. Must be set when \"l2tp-auth-method\" was selected to be \"certificate\". Insert \"defaultCert\" when you want to use the default certificate.",
									},
									"allow_vpn_clients_to_route_traffic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow VPN clients to route traffic.",
									},
									"support_nat_traversal_mechanism": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support NAT traversal mechanism (UDP encapsulation).",
									},
									"nat_traversal_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allocated NAT traversal UDP service. Identified by name or UID. Must be set when \"support-nat-traversal-mechanism\" is true.",
									},
									"support_visitor_mode": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support Visitor Mode.",
									},
									"visitor_mode_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "TCP Service for Visitor Mode. Identified by name or UID. Must be set when \"support-visitor-mode\" is true.",
									},
									"visitor_mode_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface for Visitor Mode. Must be set when \"support-visitor-mode\" is true. Insert IPV4 Address of existing interface or \"All IPs\" when you want all interfaces.",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway VPN domain identified by the name or UID.",
						},
						"vpn_domain_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway VPN domain type.",
						},
						"vpn_domain_exclude_external_ip_addresses": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Exclude the external IP addresses from the VPN domain of this Security Gateway.",
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementSimpleClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showClusterRes, err := client.ApiCall("show-simple-cluster", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showClusterRes.Success {
		return fmt.Errorf(showClusterRes.ErrorMsg)
	}

	cluster := showClusterRes.GetData()

	// If total interfaces above 50, Run show-simple-cluster with interface-limit
	if v := cluster["interfaces"]; v != nil {
		if total, ok := v.(map[string]interface{})["total"]; ok {
			totalInterfaces := int(total.(float64))
			if totalInterfaces > 50 {
				payload["limit-interfaces"] = totalInterfaces
				showClusterRes, err := client.ApiCall("show-simple-cluster", payload, client.GetSessionID(), true, client.IsProxyUsed())
				if err != nil {
					return fmt.Errorf(err.Error())
				}
				if !showClusterRes.Success {
					return fmt.Errorf(showClusterRes.ErrorMsg)
				}
				cluster = showClusterRes.GetData()
			}
		}
	}

	log.Println("Read Simple Cluster - Show JSON = ", cluster)

	if v := cluster["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := cluster["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := cluster["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := cluster["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := cluster["cluster-mode"]; v != nil {
		_ = d.Set("cluster_mode", v)
	}

	if v := cluster["geo-mode"]; v != nil {
		_ = d.Set("geo_mode", v)
	}

	if cluster["advanced-settings"] != nil {

		advancedSettingsMap, ok := cluster["advanced-settings"].(map[string]interface{})

		if ok {
			advancedSettingsMapToReturn := make(map[string]interface{})

			if v := advancedSettingsMap["connection-persistence"]; v != nil {
				advancedSettingsMapToReturn["connection_persistence"] = v
			}
			if v, ok := advancedSettingsMap["sam"]; ok {

				samMap, ok := v.(map[string]interface{})
				if ok {
					samMapToReturn := make(map[string]interface{})

					if v, _ := samMap["forward-to-other-sam-servers"]; v != nil {
						samMapToReturn["forward_to_other_sam_servers"] = v
					}
					if v, _ := samMap["use-early-versions"]; v != nil {
						samMapToReturn["use_early_versions"] = v
					}
					if v, _ := samMap["purge-sam-file"]; v != nil {
						samMapToReturn["purge_sam_file"] = v
					}
					advancedSettingsMapToReturn["sam"] = []interface{}{samMapToReturn}
				}
			}
			_ = d.Set("advanced_settings", []interface{}{advancedSettingsMapToReturn})

		}
	} else {
		_ = d.Set("advanced_settings", nil)
	}

	if v := cluster["enable-https-inspection"]; v != nil {
		_ = d.Set("enable_https_inspection", v)
	}

	if cluster["fetch-policy"] != nil {
		fetchPolicyJson, ok := cluster["fetch-policy"].([]interface{})
		if ok {
			_ = d.Set("fetch_policy", fetchPolicyJson)
		}
	} else {
		_ = d.Set("fetch_policy", nil)
	}

	if v := cluster["hit-count"]; v != nil {
		_ = d.Set("hit_count", v)
	}

	if cluster["https-inspection"] != nil {

		httpsInspectionMap, ok := cluster["https-inspection"].(map[string]interface{})

		if ok {
			httpsInspectionMapToReturn := make(map[string]interface{})

			if v, ok := httpsInspectionMap["bypass-on-failure"]; ok {

				bypassOnFailureMap, ok := v.(map[string]interface{})
				if ok {
					bypassOnFailureMapToReturn := make(map[string]interface{})

					if v, _ := bypassOnFailureMap["override-profile"]; v != nil {
						bypassOnFailureMapToReturn["override_profile"] = v
					}
					if v, _ := bypassOnFailureMap["value"]; v != nil {
						bypassOnFailureMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["bypass_on_failure"] = []interface{}{bypassOnFailureMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["site-categorization-allow-mode"]; ok {

				siteCategorizationAllowModeMap, ok := v.(map[string]interface{})
				if ok {
					siteCategorizationAllowModeMapToReturn := make(map[string]interface{})

					if v, _ := siteCategorizationAllowModeMap["override-profile"]; v != nil {
						siteCategorizationAllowModeMapToReturn["override_profile"] = v
					}
					if v, _ := siteCategorizationAllowModeMap["value"]; v != nil {
						siteCategorizationAllowModeMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["site_categorization_allow_mode"] = []interface{}{siteCategorizationAllowModeMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["deny-untrusted-server-cert"]; ok {

				denyUntrustedServerCertMap, ok := v.(map[string]interface{})
				if ok {
					denyUntrustedServerCertMapToReturn := make(map[string]interface{})

					if v, _ := denyUntrustedServerCertMap["override-profile"]; v != nil {
						denyUntrustedServerCertMapToReturn["override_profile"] = v
					}
					if v, _ := denyUntrustedServerCertMap["value"]; v != nil {
						denyUntrustedServerCertMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["deny_untrusted_server_cert"] = []interface{}{denyUntrustedServerCertMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["deny-revoked-server-cert"]; ok {

				denyRevokedServerCertMap, ok := v.(map[string]interface{})
				if ok {
					denyRevokedServerCertMapToReturn := make(map[string]interface{})

					if v, _ := denyRevokedServerCertMap["override-profile"]; v != nil {
						denyRevokedServerCertMapToReturn["override_profile"] = v
					}
					if v, _ := denyRevokedServerCertMap["value"]; v != nil {
						denyRevokedServerCertMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["deny_revoked_server_cert"] = []interface{}{denyRevokedServerCertMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["deny-expired-server-cert"]; ok {

				denyExpiredServerCertMap, ok := v.(map[string]interface{})
				if ok {
					denyExpiredServerCertMapToReturn := make(map[string]interface{})

					if v, _ := denyExpiredServerCertMap["override-profile"]; v != nil {
						denyExpiredServerCertMapToReturn["override_profile"] = v
					}
					if v, _ := denyExpiredServerCertMap["value"]; v != nil {
						denyExpiredServerCertMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["deny_expired_server_cert"] = []interface{}{denyExpiredServerCertMapToReturn}
				}
			}
			_ = d.Set("https_inspection", []interface{}{httpsInspectionMapToReturn})

		}
	} else {
		_ = d.Set("https_inspection", nil)
	}

	if v := cluster["identity-awareness"]; v != nil {
		_ = d.Set("identity_awareness", v)
	}

	if cluster["identity-awareness-settings"] != nil {

		identityAwarenessSettingsMap, ok := cluster["identity-awareness-settings"].(map[string]interface{})

		if ok {
			identityAwarenessSettingsMapToReturn := make(map[string]interface{})

			if v := identityAwarenessSettingsMap["browser-based-authentication"]; v != nil {
				identityAwarenessSettingsMapToReturn["browser_based_authentication"] = v
			}
			if v, ok := identityAwarenessSettingsMap["browser-based-authentication-settings"]; ok {

				browserBasedAuthenticationSettingsMap, ok := v.(map[string]interface{})
				if ok {
					browserBasedAuthenticationSettingsMapToReturn := make(map[string]interface{})

					if v, _ := browserBasedAuthenticationSettingsMap["authentication-settings"]; v != nil {
						browserBasedAuthenticationSettingsMapToReturn["authentication_settings"] = v
					}
					if v, _ := browserBasedAuthenticationSettingsMap["browser-based-authentication-portal-settings"]; v != nil {
						browserBasedAuthenticationSettingsMapToReturn["browser_based_authentication_portal_settings"] = v
					}
					identityAwarenessSettingsMapToReturn["browser_based_authentication_settings"] = []interface{}{browserBasedAuthenticationSettingsMapToReturn}
				}
			}
			if v := identityAwarenessSettingsMap["identity-agent"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_agent"] = v
			}
			if v, ok := identityAwarenessSettingsMap["identity-agent-settings"]; ok {

				identityAgentSettingsMap, ok := v.(map[string]interface{})
				if ok {
					identityAgentSettingsMapToReturn := make(map[string]interface{})

					if v, _ := identityAgentSettingsMap["agents-interval-keepalive"]; v != nil {
						identityAgentSettingsMapToReturn["agents_interval_keepalive"] = v
					}
					if v, _ := identityAgentSettingsMap["user-reauthenticate-interval"]; v != nil {
						identityAgentSettingsMapToReturn["user_reauthenticate_interval"] = v
					}
					if v, _ := identityAgentSettingsMap["authentication-settings"]; v != nil {
						identityAgentSettingsMapToReturn["authentication_settings"] = v
					}
					if v, _ := identityAgentSettingsMap["identity-agent-portal-settings"]; v != nil {
						identityAgentSettingsMapToReturn["identity_agent_portal_settings"] = v
					}
					identityAwarenessSettingsMapToReturn["identity_agent_settings"] = []interface{}{identityAgentSettingsMapToReturn}
				}
			}
			if v := identityAwarenessSettingsMap["identity-collector"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_collector"] = v
			}
			if v, ok := identityAwarenessSettingsMap["identity-collector-settings"]; ok {

				identityCollectorSettingsMap, ok := v.(map[string]interface{})
				if ok {
					identityCollectorSettingsMapToReturn := make(map[string]interface{})

					if v, _ := identityCollectorSettingsMap["authorized-clients"]; v != nil {
						identityCollectorSettingsMapToReturn["authorized_clients"] = v
					}
					if v, _ := identityCollectorSettingsMap["authentication-settings"]; v != nil {
						identityCollectorSettingsMapToReturn["authentication_settings"] = v
					}
					if v, _ := identityCollectorSettingsMap["client-access-permissions"]; v != nil {
						identityCollectorSettingsMapToReturn["client_access_permissions"] = v
					}
					identityAwarenessSettingsMapToReturn["identity_collector_settings"] = []interface{}{identityCollectorSettingsMapToReturn}
				}
			}
			if v, ok := identityAwarenessSettingsMap["identity-sharing-settings"]; ok {

				identitySharingSettingsMap, ok := v.(map[string]interface{})
				if ok {
					identitySharingSettingsMapToReturn := make(map[string]interface{})

					if v, _ := identitySharingSettingsMap["share-with-other-gateways"]; v != nil {
						identitySharingSettingsMapToReturn["share_with_other_gateways"] = v
					}
					if v, _ := identitySharingSettingsMap["receive-from-other-gateways"]; v != nil {
						identitySharingSettingsMapToReturn["receive_from_other_gateways"] = v
					}
					if v, _ := identitySharingSettingsMap["receive-from"]; v != nil {
						identitySharingSettingsMapToReturn["receive_from"] = v
					}
					identityAwarenessSettingsMapToReturn["identity_sharing_settings"] = []interface{}{identitySharingSettingsMapToReturn}
				}
			}
			if v, ok := identityAwarenessSettingsMap["proxy-settings"]; ok {

				proxySettingsMap, ok := v.(map[string]interface{})
				if ok {
					proxySettingsMapToReturn := make(map[string]interface{})

					if v, _ := proxySettingsMap["detect-using-x-forward-for"]; v != nil {
						proxySettingsMapToReturn["detect_using_x_forward_for"] = v
					}
					identityAwarenessSettingsMapToReturn["proxy_settings"] = []interface{}{proxySettingsMapToReturn}
				}
			}
			if v := identityAwarenessSettingsMap["remote-access"]; v != nil {
				identityAwarenessSettingsMapToReturn["remote_access"] = v
			}
			_ = d.Set("identity_awareness_settings", []interface{}{identityAwarenessSettingsMapToReturn})

		}
	} else {
		_ = d.Set("identity_awareness_settings", nil)
	}

	if v := cluster["ips-update-policy"]; v != nil {
		_ = d.Set("ips_update_policy", v)
	}

	if v := cluster["nat-hide-internal-interfaces"]; v != nil {
		_ = d.Set("nat_hide_internal_interfaces", v)
	}

	if cluster["nat-settings"] != nil {

		natSettingsMap := cluster["nat-settings"].(map[string]interface{})

		natSettingsMapToReturn := make(map[string]interface{})

		if v := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = v
		}
		if v := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}
		if v := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}
		if v := natSettingsMap["ipv4-address"]; v != nil {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v := natSettingsMap["ipv6-address"]; v != nil {
			natSettingsMapToReturn["ipv6_address"] = v
		}
		if v := natSettingsMap["method"]; v != nil {
			natSettingsMapToReturn["method"] = v
		}

		_ = d.Set("nat_settings", []interface{}{natSettingsMapToReturn})

	} else {
		_ = d.Set("nat_settings", nil)
	}

	if cluster["platform-portal-settings"] != nil {

		platformPortalSettingsMap, ok := cluster["platform-portal-settings"].(map[string]interface{})

		if ok {
			platformPortalSettingsMapToReturn := make(map[string]interface{})

			if v, ok := platformPortalSettingsMap["portal-web-settings"]; ok {

				portalWebSettingsMap, ok := v.(map[string]interface{})
				if ok {
					portalWebSettingsMapToReturn := make(map[string]interface{})

					if v, _ := portalWebSettingsMap["aliases"]; v != nil {
						portalWebSettingsMapToReturn["aliases"] = v
					}
					if v, _ := portalWebSettingsMap["main-url"]; v != nil {
						portalWebSettingsMapToReturn["main_url"] = v
					}
					platformPortalSettingsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
				}
			}
			if v, ok := platformPortalSettingsMap["certificate-settings"]; ok {

				certificateSettingsMap, ok := v.(map[string]interface{})
				if ok {
					certificateSettingsMapToReturn := make(map[string]interface{})

					if v, _ := certificateSettingsMap["base64-certificate"]; v != nil {
						certificateSettingsMapToReturn["base64_certificate"] = v
					}
					if v, _ := certificateSettingsMap["base64-password"]; v != nil {
						certificateSettingsMapToReturn["base64_password"] = v
					}
					platformPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
				}
			}
			if v, ok := platformPortalSettingsMap["accessibility"]; ok {

				accessibilityMap, ok := v.(map[string]interface{})
				if ok {
					accessibilityMapToReturn := make(map[string]interface{})

					if v, _ := accessibilityMap["allow-access-from"]; v != nil {
						accessibilityMapToReturn["allow_access_from"] = v
					}
					if v, _ := accessibilityMap["internal-access-settings"]; v != nil {
						accessibilityMapToReturn["internal_access_settings"] = v
					}
					platformPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
				}
			}
			_ = d.Set("platform_portal_settings", []interface{}{platformPortalSettingsMapToReturn})

		}
	} else {
		_ = d.Set("platform_portal_settings", nil)
	}

	if cluster["proxy-settings"] != nil {

		proxySettingsMap := cluster["proxy-settings"].(map[string]interface{})

		proxySettingsMapToReturn := make(map[string]interface{})

		if v := proxySettingsMap["use-custom-proxy"]; v != nil {
			proxySettingsMapToReturn["use_custom_proxy"] = v
		}
		if v := proxySettingsMap["proxy-server"]; v != nil {
			proxySettingsMapToReturn["proxy_server"] = v
		}
		if v := proxySettingsMap["port"]; v != nil {
			proxySettingsMapToReturn["port"] = v
		}

		_ = d.Set("proxy_settings", []interface{}{proxySettingsMapToReturn})

	} else {
		_ = d.Set("proxy_settings", nil)
	}

	if v := cluster["qos"]; v != nil {
		_ = d.Set("qos", v)
	}

	if cluster["usercheck-portal-settings"] != nil {

		usercheckPortalSettingsMap, ok := cluster["usercheck-portal-settings"].(map[string]interface{})

		if ok {
			usercheckPortalSettingsMapToReturn := make(map[string]interface{})

			if v := usercheckPortalSettingsMap["enabled"]; v != nil {
				usercheckPortalSettingsMapToReturn["enabled"] = v
			}
			if v, ok := usercheckPortalSettingsMap["portal-web-settings"]; ok {

				portalWebSettingsMap, ok := v.(map[string]interface{})
				if ok {
					portalWebSettingsMapToReturn := make(map[string]interface{})

					if v, _ := portalWebSettingsMap["aliases"]; v != nil {
						portalWebSettingsMapToReturn["aliases"] = v
					}
					if v, _ := portalWebSettingsMap["main-url"]; v != nil {
						portalWebSettingsMapToReturn["main_url"] = v
					}
					usercheckPortalSettingsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
				}
			}
			if v, ok := usercheckPortalSettingsMap["certificate-settings"]; ok {

				certificateSettingsMap, ok := v.(map[string]interface{})
				if ok {
					certificateSettingsMapToReturn := make(map[string]interface{})

					if v, _ := certificateSettingsMap["base64-certificate"]; v != nil {
						certificateSettingsMapToReturn["base64_certificate"] = v
					}
					if v, _ := certificateSettingsMap["base64-password"]; v != nil {
						certificateSettingsMapToReturn["base64_password"] = v
					}
					usercheckPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
				}
			}
			if v, ok := usercheckPortalSettingsMap["accessibility"]; ok {

				accessibilityMap, ok := v.(map[string]interface{})
				if ok {
					accessibilityMapToReturn := make(map[string]interface{})

					if v, _ := accessibilityMap["allow-access-from"]; v != nil {
						accessibilityMapToReturn["allow_access_from"] = v
					}
					if v, _ := accessibilityMap["internal-access-settings"]; v != nil {
						accessibilityMapToReturn["internal_access_settings"] = v
					}
					usercheckPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
				}
			}
			_ = d.Set("usercheck_portal_settings", []interface{}{usercheckPortalSettingsMapToReturn})

		}
	} else {
		_ = d.Set("usercheck_portal_settings", nil)
	}

	if v := cluster["zero-phishing"]; v != nil {
		_ = d.Set("zero_phishing", v)
	}

	if v := cluster["zero-phishing-fqdn"]; v != nil {
		_ = d.Set("zero_phishing_fqdn", v)
	}

	if v := cluster["interfaces"]; v != nil {
		interfacesList := v.(map[string]interface{})["objects"].([]interface{})
		if len(interfacesList) > 0 {
			var interfacesListState []map[string]interface{}
			for i := range interfacesList {
				interfaceJson := interfacesList[i].(map[string]interface{})
				interfaceState := make(map[string]interface{})
				if v, _ := interfaceJson["name"]; v != nil {
					interfaceState["name"] = v
				}
				if v, _ := interfaceJson["ipv4-address"]; v != nil {
					interfaceState["ipv4_address"] = v
				}
				if v, _ := interfaceJson["ipv4-mask-length"]; v != nil {
					interfaceState["ipv4_mask_length"] = v
				}
				if v, _ := interfaceJson["ipv4-network-mask"]; v != nil {
					interfaceState["ipv4_network_mask"] = v
				}
				if v, _ := interfaceJson["ipv6-address"]; v != nil {
					interfaceState["ipv6_address"] = v
				}
				if v, _ := interfaceJson["ipv6-mask-length"]; v != nil {
					interfaceState["ipv6_mask_length"] = v
				}
				if v, _ := interfaceJson["ipv6-network-mask"]; v != nil {
					interfaceState["ipv6_network_mask"] = v
				}
				if v, _ := interfaceJson["interface-type"]; v != nil {
					interfaceState["interface_type"] = v
				}
				if v, _ := interfaceJson["anti-spoofing"]; v != nil {
					interfaceState["anti_spoofing"] = v
				}
				if v, _ := interfaceJson["anti-spoofing-settings"]; v != nil {
					antiSpoofingSettingsJson := v.(map[string]interface{})
					antiSpoofingSettingsState := make(map[string]interface{})
					if v, _ := antiSpoofingSettingsJson["action"]; v != nil {
						antiSpoofingSettingsState["action"] = v
					}
					interfaceState["anti_spoofing_settings"] = antiSpoofingSettingsState
				}
				if v, _ := interfaceJson["security-zone"]; v != nil {
					interfaceState["security_zone"] = v
				}
				if v, _ := interfaceJson["security-zone-settings"]; v != nil {
					securityZoneSettingsJson := v.(map[string]interface{})
					securityZoneSettingsState := make(map[string]interface{})
					if v, _ := securityZoneSettingsJson["auto-calculated"]; v != nil {
						securityZoneSettingsState["auto_calculated"] = v
					}
					if v, _ := securityZoneSettingsJson["specific-zone"]; v != nil {
						securityZoneSettingsState["specific_zone"] = v
					}
					interfaceState["security_zone_settings"] = securityZoneSettingsState
				}
				if v, _ := interfaceJson["topology"]; v != nil {
					interfaceState["topology"] = v
				}
				if v, _ := interfaceJson["topology-automatic-calculation"]; v != nil {
					interfaceState["topology_automatic_calculation"] = v
				}
				if v, _ := interfaceJson["topology-settings"]; v != nil {
					topologySettingsJson := v.(map[string]interface{})
					topologySettingsState := make(map[string]interface{})
					if v, _ := topologySettingsJson["interface-leads-to-dmz"]; v != nil {
						topologySettingsState["interface_leads_to_dmz"] = v
					}
					if v, _ := topologySettingsJson["ip-address-behind-this-interface"]; v != nil {
						topologySettingsState["ip_address_behind_this_interface"] = v
					}
					if v, _ := topologySettingsJson["specific-network"]; v != nil {
						topologySettingsState["specific_network"] = v
					}
					interfaceState["topology_settings"] = topologySettingsState
				}

				if v, _ := interfaceJson["color"]; v != nil {
					interfaceState["color"] = v
				}
				if v, _ := interfaceJson["comments"]; v != nil {
					interfaceState["comments"] = v
				}
				interfacesListState = append(interfacesListState, interfaceState)
			}
			_ = d.Set("interfaces", interfacesListState)
		} else {
			_ = d.Set("interfaces", interfacesList)
		}
	} else {
		_ = d.Set("interfaces", nil)
	}

	if v := cluster["cluster-members"]; v != nil {
		membersList := v.([]interface{})
		if len(membersList) > 0 {
			var membersListState []map[string]interface{}
			for i := range membersList {
				memberJson := membersList[i].(map[string]interface{})
				memberState := make(map[string]interface{})
				if v, _ := memberJson["name"]; v != nil {
					memberState["name"] = v
				}
				if v, _ := memberJson["ip-address"]; v != nil {
					memberState["ip_address"] = v
				}
				if v, _ := memberJson["interfaces"]; v != nil {
					memberInterfacesList := v.([]interface{})
					if len(memberInterfacesList) > 0 {
						var memberInterfacesState []map[string]interface{}
						for i := range memberInterfacesList {
							memberInterfaceJson := memberInterfacesList[i].(map[string]interface{})
							memberInterfaceState := make(map[string]interface{})
							if v, _ := memberInterfaceJson["name"]; v != nil {
								memberInterfaceState["name"] = v
							}
							if v, _ := memberInterfaceJson["ipv4-address"]; v != nil {
								memberInterfaceState["ipv4_address"] = v
							}
							if v, _ := memberInterfaceJson["ipv4-mask-length"]; v != nil {
								memberInterfaceState["ipv4_mask_length"] = v
							}
							if v, _ := memberInterfaceJson["ipv4-network-mask"]; v != nil {
								memberInterfaceState["ipv4_network_mask"] = v
							}
							if v, _ := memberInterfaceJson["ipv6-address"]; v != nil {
								memberInterfaceState["ipv6_address"] = v
							}
							if v, _ := memberInterfaceJson["ipv6-mask-length"]; v != nil {
								memberInterfaceState["ipv6_mask_length"] = v
							}
							if v, _ := memberInterfaceJson["ipv6-network-mask"]; v != nil {
								memberInterfaceState["ipv6_network_mask"] = v
							}
							memberInterfacesState = append(memberInterfacesState, memberInterfaceState)
						}
						memberState["interfaces"] = memberInterfacesState
					}
				}

				if v, _ := memberJson["sic-message"]; v != nil {
					memberState["sic_message"] = v
				}
				if v, _ := memberJson["sic-state"]; v != nil {
					memberState["sic_state"] = v
				}
				membersListState = append(membersListState, memberState)
			}
			_ = d.Set("members", membersListState)
		} else {
			_ = d.Set("members", membersList)
		}
	} else {
		_ = d.Set("members", nil)
	}

	if v := cluster["anti-bot"]; v != nil {
		_ = d.Set("anti_bot", v)
	}

	if v := cluster["anti-virus"]; v != nil {
		_ = d.Set("anti_virus", v)
	}

	if v := cluster["application-control"]; v != nil {
		_ = d.Set("application_control", v)
	}

	if v := cluster["content-awareness"]; v != nil {
		_ = d.Set("content_awareness", v)
	}

	if v := cluster["dynamic-ip"]; v != nil {
		_ = d.Set("dynamic_ip", v)
	}

	if v := cluster["firewall"]; v != nil {
		_ = d.Set("firewall", v)
	}

	if v := cluster["ips"]; v != nil {
		_ = d.Set("ips", v)
	}

	if cluster["ips-settings"] != nil {

		ipsSettingsMap, ok := cluster["ips-settings"].(map[string]interface{})

		if ok {
			ipsSettingsMapToReturn := make(map[string]interface{})

			if v := ipsSettingsMap["bypass-all-under-load"]; v != nil {
				ipsSettingsMapToReturn["bypass_all_under_load"] = v
			}
			if v := ipsSettingsMap["bypass-track-method"]; v != nil {
				ipsSettingsMapToReturn["bypass_track_method"] = v
			}
			if v, ok := ipsSettingsMap["top-cpu-consuming-protections"]; ok {

				topCpuConsumingProtectionsMap, ok := v.(map[string]interface{})
				if ok {
					topCpuConsumingProtectionsMapToReturn := make(map[string]interface{})

					if v, _ := topCpuConsumingProtectionsMap["disable-period"]; v != nil {
						topCpuConsumingProtectionsMapToReturn["disable_period"] = v
					}
					if v, _ := topCpuConsumingProtectionsMap["disable-under-load"]; v != nil {
						topCpuConsumingProtectionsMapToReturn["disable_under_load"] = v
					}
					ipsSettingsMapToReturn["top_cpu_consuming_protections"] = []interface{}{topCpuConsumingProtectionsMapToReturn}
				}
			}
			if v := ipsSettingsMap["activation-mode"]; v != nil {
				ipsSettingsMapToReturn["activation_mode"] = v
			}
			if v := ipsSettingsMap["cpu-usage-low-threshold"]; v != nil {
				ipsSettingsMapToReturn["cpu_usage_low_threshold"] = v
			}
			if v := ipsSettingsMap["cpu-usage-high-threshold"]; v != nil {
				ipsSettingsMapToReturn["cpu_usage_high_threshold"] = v
			}
			if v := ipsSettingsMap["memory-usage-low-threshold"]; v != nil {
				ipsSettingsMapToReturn["memory_usage_low_threshold"] = v
			}
			if v := ipsSettingsMap["memory-usage-high-threshold"]; v != nil {
				ipsSettingsMapToReturn["memory_usage_high_threshold"] = v
			}
			if v := ipsSettingsMap["send-threat-cloud-info"]; v != nil {
				ipsSettingsMapToReturn["send_threat_cloud_info"] = v
			}
			if v := ipsSettingsMap["reject-on-cluster-fail-over"]; v != nil {
				ipsSettingsMapToReturn["reject_on_cluster_fail_over"] = v
			}
			_ = d.Set("ips_settings", []interface{}{ipsSettingsMapToReturn})

		}
	} else {
		_ = d.Set("ips_settings", nil)
	}

	if v := cluster["threat-emulation"]; v != nil {
		_ = d.Set("threat_emulation", v)
	}

	if v := cluster["url-filtering"]; v != nil {
		_ = d.Set("url_filtering", v)
	}

	if v := cluster["vpn"]; v != nil {
		_ = d.Set("vpn", v)
	}

	if v := cluster["os-name"]; v != nil {
		_ = d.Set("os_name", v)
	}

	if v := cluster["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := cluster["hardware"]; v != nil {
		_ = d.Set("hardware", v)
	}

	if v := cluster["sic-name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	if v := cluster["sic-state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if v := cluster["save-logs-locally"]; v != nil {
		_ = d.Set("save_logs_locally", v)
	}

	if v := cluster["send_alerts_to_server"]; v != nil {
		_ = d.Set("send_alerts_to_server", v)
	} else {
		_ = d.Set("send_alerts_to_server", nil)
	}

	if v := cluster["send-logs-to-backup-server"]; v != nil {
		_ = d.Set("send_logs_to_backup_server", v)
	} else {
		_ = d.Set("send_logs_to_backup_server", nil)
	}

	if v := cluster["send-logs-to-server"]; v != nil {
		_ = d.Set("send_logs_to_server", v)
	} else {
		_ = d.Set("send_logs_to_server", nil)
	}

	if cluster["logs-settings"] != nil {

		logsSettingsMap := cluster["logs-settings"].(map[string]interface{})

		logsSettingsMapToReturn := make(map[string]interface{})

		if v := logsSettingsMap["alert-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-type"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_type"] = v
		}
		if v := logsSettingsMap["before-delete-keep-logs-from-the-last-days"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days"] = v
		}
		if v := logsSettingsMap["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days_threshold"] = v
		}
		if v := logsSettingsMap["before-delete-run-script"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script"] = v
		}
		if v := logsSettingsMap["before-delete-run-script-command"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script_command"] = v
		}
		if v := logsSettingsMap["delete-index-files-older-than-days"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days"] = v
		}
		if v := logsSettingsMap["delete-index-files-older-than-days-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days_threshold"] = v
		}
		if v := logsSettingsMap["delete-index-files-when-index-size-above"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_when_index_size_above"] = v
		}
		if v := logsSettingsMap["delete-index-files-when-index-size-above-metrics"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_when_index_size_above_metrics"] = v
		}
		if v := logsSettingsMap["delete-index-files-when-index-size-above-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_when_index_size_above_threshold"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["detect-new-citrix-ica-application-names"]; v != nil {
			logsSettingsMapToReturn["detect_new_citrix_ica_application_names"] = v
		}
		if v := logsSettingsMap["distribute-logs-between-all-active-servers"]; v != nil {
			logsSettingsMapToReturn["distribute_logs_between_all_active_servers"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_name"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server-schedule-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_schedule_name"] = v
		}
		if v := logsSettingsMap["perform-log-rotate-before-log-forwarding"]; v != nil {
			logsSettingsMapToReturn["perform_log_rotate_before_log_forwarding"] = v
		}
		if v := logsSettingsMap["reject-connections-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["reject_connections_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["reserve-for-packet-capture-metrics"]; v != nil {
			logsSettingsMapToReturn["reserve_for_packet_capture_metrics"] = v
		}
		if v := logsSettingsMap["reserve-for-packet-capture-threshold"]; v != nil {
			logsSettingsMapToReturn["reserve_for_packet_capture_threshold"] = v
		}
		if v := logsSettingsMap["rotate-log-by-file-size"]; v != nil {
			logsSettingsMapToReturn["rotate_log_by_file_size"] = v
		}
		if v := logsSettingsMap["rotate-log-file-size-threshold"]; v != nil {
			logsSettingsMapToReturn["rotate_log_file_size_threshold"] = v
		}
		if v := logsSettingsMap["rotate-log-on-schedule"]; v != nil {
			logsSettingsMapToReturn["rotate_log_on_schedule"] = v
		}
		if v := logsSettingsMap["rotate-log-schedule-name"]; v != nil {
			logsSettingsMapToReturn["rotate_log_schedule_name"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["turn-on-qos-logging"]; v != nil {
			logsSettingsMapToReturn["turn_on_qos_logging"] = v
		}
		if v := logsSettingsMap["update-account-log-every"]; v != nil {
			logsSettingsMapToReturn["update_account_log_every"] = v
		}

		_ = d.Set("logs_settings", []interface{}{logsSettingsMapToReturn})

	} else {
		_ = d.Set("logs_settings", nil)
	}

	if cluster["firewall-settings"] != nil {

		firewallSettingsMap := cluster["firewall-settings"].(map[string]interface{})

		firewallSettingsMapToReturn := make(map[string]interface{})

		if v := firewallSettingsMap["auto-calculate-connections-hash-table-size-and-memory-pool"]; v != nil {
			firewallSettingsMapToReturn["auto_calculate_connections_hash_table_size_and_memory_pool"] = v
		}
		if v := firewallSettingsMap["auto-maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsMapToReturn["auto_maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsMap["connections-hash-size"]; v != nil {
			firewallSettingsMapToReturn["connections_hash_size"] = v
		}
		if v := firewallSettingsMap["maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsMapToReturn["maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsMap["maximum-memory-pool-size"]; v != nil {
			firewallSettingsMapToReturn["maximum_memory_pool_size"] = v
		}
		if v := firewallSettingsMap["memory-pool-size"]; v != nil {
			firewallSettingsMapToReturn["memory_pool_size"] = v
		}

		_ = d.Set("firewall_settings", []interface{}{firewallSettingsMapToReturn})

	} else {
		_ = d.Set("firewall_settings", nil)
	}

	if cluster["vpn-settings"] != nil {

		vpnSettingsMap := cluster["vpn-settings"].(map[string]interface{})

		vpnSettingsMapToReturn := make(map[string]interface{})

		if v := vpnSettingsMap["interfaces"]; v != nil {

			interfacesList := v.([]interface{})

			if len(interfacesList) > 0 {

				var interfacesListToReturn []map[string]interface{}

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v := interfacesMap["interface-name"]; v != nil {
						interfacesMapToAdd["interface_name"] = v
					}
					if v := interfacesMap["next-hop-ip"]; v != nil {
						interfacesMapToAdd["next_hop_ip"] = v
					}
					if v := interfacesMap["static-nat-ip"]; v != nil {
						interfacesMapToAdd["static_nat_ip"] = v
					}
					if v := interfacesMap["priority"]; v != nil {
						interfacesMapToAdd["priority"] = v
					}
					if v := interfacesMap["redundancy-mode"]; v != nil {
						interfacesMapToAdd["redundancy_mode"] = v
					}
					if v := interfacesMap["ip-version"]; v != nil {
						interfacesMapToAdd["ip_version"] = v
					}

					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}

				vpnSettingsMapToReturn["interfaces"] = interfacesListToReturn
			}
		}

		if v := vpnSettingsMap["advanced"]; v != nil {

			advancedMap := v.(map[string]interface{})

			advancedMapToReturn := make(map[string]interface{})

			if v := advancedMap["tunnel-sharing-mode"]; v != nil {
				advancedMapToReturn["tunnel_sharing_mode"] = v
			}
			if v := advancedMap["shutdown-on-gateway-restart"]; v != nil {
				advancedMapToReturn["shutdown_on_gateway_restart"] = v
			}
			if v := advancedMap["enable-wire-mode"]; v != nil {
				advancedMapToReturn["enable_wire_mode"] = v
			}
			if v := advancedMap["wire-mode-interfaces"]; v != nil {

				wireModeInterfacesList := v.([]interface{})

				if len(wireModeInterfacesList) > 0 {

					var wireModeInterfacesListToReturn []map[string]interface{}

					for i := range wireModeInterfacesList {

						wireModeInterfacesMap := wireModeInterfacesList[i].(map[string]interface{})

						wireModeInterfacesMapToAdd := make(map[string]interface{})

						if v := wireModeInterfacesMap["name"]; v != nil {
							wireModeInterfacesMapToAdd["name"] = v
						}
						if v := wireModeInterfacesMap["ip-address"]; v != nil {
							wireModeInterfacesMapToAdd["ip_address"] = v
						}
						if v := wireModeInterfacesMap["netmask"]; v != nil {
							wireModeInterfacesMapToAdd["netmask"] = v
						}

						wireModeInterfacesListToReturn = append(wireModeInterfacesListToReturn, wireModeInterfacesMapToAdd)
					}

					advancedMapToReturn["wire_mode_interfaces"] = wireModeInterfacesListToReturn
				}
			}

			if v := advancedMap["enable-wire-mode-log-traffic"]; v != nil {
				advancedMapToReturn["enable_wire_mode_log_traffic"] = v
			}
			if v := advancedMap["enable-nat-traversal"]; v != nil {
				advancedMapToReturn["enable_nat_traversal"] = v
			}

			vpnSettingsMapToReturn["advanced"] = []interface{}{advancedMapToReturn}
		}

		if v := vpnSettingsMap["authentication"]; v != nil {

			authenticationMap := v.(map[string]interface{})

			authenticationMapToReturn := make(map[string]interface{})

			if v := authenticationMap["authentication-clients"]; v != nil {

				authenticationClientsList := v.([]interface{})

				if len(authenticationClientsList) > 0 {

					var authenticationClientsListToReturn []map[string]interface{}

					for i := range authenticationClientsList {

						authenticationClientsMap := authenticationClientsList[i].(map[string]interface{})

						authenticationClientsMapToAdd := make(map[string]interface{})

						if v := authenticationClientsMap["name"]; v != nil {
							authenticationClientsMapToAdd["name"] = v
						}
						if v := authenticationClientsMap["type"]; v != nil {
							authenticationClientsMapToAdd["type"] = v
						}
						if v := authenticationClientsMap["color"]; v != nil {
							authenticationClientsMapToAdd["color"] = v
						}
						if v := authenticationClientsMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							authenticationClientsMapToAdd["domain"] = []interface{}{domainMapToReturn}
						}

						if v := authenticationClientsMap["icon"]; v != nil {
							authenticationClientsMapToAdd["icon"] = v
						}
						if v := authenticationClientsMap["uid"]; v != nil {
							authenticationClientsMapToAdd["uid"] = v
						}

						authenticationClientsListToReturn = append(authenticationClientsListToReturn, authenticationClientsMapToAdd)
					}

					authenticationMapToReturn["authentication_clients"] = authenticationClientsListToReturn
				}
			}

			vpnSettingsMapToReturn["authentication"] = []interface{}{authenticationMapToReturn}
		}

		if v := vpnSettingsMap["certificates"]; v != nil {

			certificatesList := v.([]interface{})

			if len(certificatesList) > 0 {

				var certificatesListToReturn []map[string]interface{}

				for i := range certificatesList {

					certificatesMap := certificatesList[i].(map[string]interface{})

					certificatesMapToAdd := make(map[string]interface{})

					if v := certificatesMap["name"]; v != nil {
						certificatesMapToAdd["name"] = v
					}
					if v := certificatesMap["distinguished-name"]; v != nil {
						certificatesMapToAdd["distinguished_name"] = v
					}
					if v := certificatesMap["type"]; v != nil {
						certificatesMapToAdd["type"] = v
					}
					if v := certificatesMap["base64-certificate"]; v != nil {
						certificatesMapToAdd["base64_certificate"] = v
					}
					if v := certificatesMap["certificate-authority"]; v != nil {

						certificateAuthorityMap := v.(map[string]interface{})

						certificateAuthorityMapToReturn := make(map[string]interface{})

						if v := certificateAuthorityMap["name"]; v != nil {
							certificateAuthorityMapToReturn["name"] = v
						}
						if v := certificateAuthorityMap["type"]; v != nil {
							certificateAuthorityMapToReturn["type"] = v
						}
						if v := certificateAuthorityMap["color"]; v != nil {
							certificateAuthorityMapToReturn["color"] = v
						}
						if v := certificateAuthorityMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							certificateAuthorityMapToReturn["domain"] = []interface{}{domainMapToReturn}
						}

						if v := certificateAuthorityMap["icon"]; v != nil {
							certificateAuthorityMapToReturn["icon"] = v
						}
						if v := certificateAuthorityMap["uid"]; v != nil {
							certificateAuthorityMapToReturn["uid"] = v
						}

						certificatesMapToAdd["certificate_authority"] = []interface{}{certificateAuthorityMapToReturn}
					}

					if v := certificatesMap["expiration-date"]; v != nil {

						expirationDateMap := v.(map[string]interface{})

						expirationDateMapToReturn := make(map[string]interface{})

						if v := expirationDateMap["iso-8601"]; v != nil {
							expirationDateMapToReturn["iso_8601"] = v
						}
						if v := expirationDateMap["posix"]; v != nil {
							expirationDateMapToReturn["posix"] = v
						}

						certificatesMapToAdd["expiration_date"] = []interface{}{expirationDateMapToReturn}
					}

					if v := certificatesMap["status"]; v != nil {
						certificatesMapToAdd["status"] = v
					}
					if v := certificatesMap["stored-at"]; v != nil {
						certificatesMapToAdd["stored_at"] = v
					}
					if v := certificatesMap["domain"]; v != nil {

						domainMap := v.(map[string]interface{})

						domainMapToReturn := make(map[string]interface{})

						if v := domainMap["name"]; v != nil {
							domainMapToReturn["name"] = v
						}
						if v := domainMap["domain-type"]; v != nil {
							domainMapToReturn["domain_type"] = v
						}
						if v := domainMap["uid"]; v != nil {
							domainMapToReturn["uid"] = v
						}

						certificatesMapToAdd["domain"] = []interface{}{domainMapToReturn}
					}

					certificatesListToReturn = append(certificatesListToReturn, certificatesMapToAdd)
				}

				vpnSettingsMapToReturn["certificates"] = certificatesListToReturn
			}
		}

		if v := vpnSettingsMap["exported-routes"]; v != nil {

			exportedRoutesMap := v.(map[string]interface{})

			exportedRoutesMapToReturn := make(map[string]interface{})

			if v := exportedRoutesMap["internal-interfaces"]; v != nil {
				exportedRoutesMapToReturn["internal_interfaces"] = v
			}
			if v := exportedRoutesMap["static-routes"]; v != nil {
				exportedRoutesMapToReturn["static_routes"] = v
			}
			if v := exportedRoutesMap["custom-routes"]; v != nil {
				exportedRoutesMapToReturn["custom_routes"] = v
			}
			if v := exportedRoutesMap["custom-routes-object"]; v != nil {

				customRoutesObjectMap := v.(map[string]interface{})

				customRoutesObjectMapToReturn := make(map[string]interface{})

				if v := customRoutesObjectMap["name"]; v != nil {
					customRoutesObjectMapToReturn["name"] = v
				}
				if v := customRoutesObjectMap["type"]; v != nil {
					customRoutesObjectMapToReturn["type"] = v
				}
				if v := customRoutesObjectMap["color"]; v != nil {
					customRoutesObjectMapToReturn["color"] = v
				}
				if v := customRoutesObjectMap["domain"]; v != nil {

					domainMap := v.(map[string]interface{})

					domainMapToReturn := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToReturn["name"] = v
					}
					if v := domainMap["domain-type"]; v != nil {
						domainMapToReturn["domain_type"] = v
					}
					if v := domainMap["uid"]; v != nil {
						domainMapToReturn["uid"] = v
					}

					customRoutesObjectMapToReturn["domain"] = []interface{}{domainMapToReturn}
				}

				if v := customRoutesObjectMap["icon"]; v != nil {
					customRoutesObjectMapToReturn["icon"] = v
				}
				if v := customRoutesObjectMap["uid"]; v != nil {
					customRoutesObjectMapToReturn["uid"] = v
				}

				exportedRoutesMapToReturn["custom_routes_object"] = []interface{}{customRoutesObjectMapToReturn}
			}

			vpnSettingsMapToReturn["exported_routes"] = []interface{}{exportedRoutesMapToReturn}
		}

		if v := vpnSettingsMap["link-selection"]; v != nil {

			linkSelectionMap := v.(map[string]interface{})

			linkSelectionMapToReturn := make(map[string]interface{})

			if v := linkSelectionMap["ip-selection"]; v != nil {
				linkSelectionMapToReturn["ip_selection"] = v
			}
			if v := linkSelectionMap["ip-address"]; v != nil {
				linkSelectionMapToReturn["ip_address"] = v
			}
			if v := linkSelectionMap["dns-resolving-hostname"]; v != nil {
				linkSelectionMapToReturn["dns_resolving_hostname"] = v
			}
			if v := linkSelectionMap["route-selection-method"]; v != nil {
				linkSelectionMapToReturn["route_selection_method"] = v
			}
			if v := linkSelectionMap["responding-traffic"]; v != nil {
				linkSelectionMapToReturn["responding_traffic"] = v
			}
			if v := linkSelectionMap["source-ip-selection"]; v != nil {
				linkSelectionMapToReturn["source_ip_selection"] = v
			}
			if v := linkSelectionMap["selected-ip"]; v != nil {
				linkSelectionMapToReturn["selected_ip"] = v
			}
			if v := linkSelectionMap["outgoing-link-tracking"]; v != nil {
				linkSelectionMapToReturn["outgoing_link_tracking"] = v
			}

			vpnSettingsMapToReturn["link_selection"] = []interface{}{linkSelectionMapToReturn}
		}

		if v := vpnSettingsMap["maximum-concurrent-ike-negotiations"]; v != nil {
			vpnSettingsMapToReturn["maximum_concurrent_ike_negotiations"] = v
		}
		if v := vpnSettingsMap["maximum-concurrent-tunnels"]; v != nil {
			vpnSettingsMapToReturn["maximum_concurrent_tunnels"] = v
		}
		if v := vpnSettingsMap["office-mode"]; v != nil {

			officeModeMap := v.(map[string]interface{})

			officeModeMapToReturn := make(map[string]interface{})

			if v := officeModeMap["mode"]; v != nil {
				officeModeMapToReturn["mode"] = v
			}
			if v := officeModeMap["group"]; v != nil {

				groupMap := v.(map[string]interface{})

				groupMapToReturn := make(map[string]interface{})

				if v := groupMap["name"]; v != nil {
					groupMapToReturn["name"] = v
				}
				if v := groupMap["type"]; v != nil {
					groupMapToReturn["type"] = v
				}
				if v := groupMap["color"]; v != nil {
					groupMapToReturn["color"] = v
				}
				if v := groupMap["domain"]; v != nil {

					domainMap := v.(map[string]interface{})

					domainMapToReturn := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToReturn["name"] = v
					}
					if v := domainMap["domain-type"]; v != nil {
						domainMapToReturn["domain_type"] = v
					}
					if v := domainMap["uid"]; v != nil {
						domainMapToReturn["uid"] = v
					}

					groupMapToReturn["domain"] = []interface{}{domainMapToReturn}
				}

				if v := groupMap["icon"]; v != nil {
					groupMapToReturn["icon"] = v
				}
				if v := groupMap["uid"]; v != nil {
					groupMapToReturn["uid"] = v
				}

				officeModeMapToReturn["group"] = []interface{}{groupMapToReturn}
			}

			if v := officeModeMap["allocate-ip-address-from"]; v != nil {

				allocateIpAddressFromMap := v.(map[string]interface{})

				allocateIpAddressFromMapToReturn := make(map[string]interface{})

				if v := allocateIpAddressFromMap["radius-server"]; v != nil {
					allocateIpAddressFromMapToReturn["radius_server"] = v
				}
				if v := allocateIpAddressFromMap["use-allocate-method"]; v != nil {
					allocateIpAddressFromMapToReturn["use_allocate_method"] = v
				}
				if v := allocateIpAddressFromMap["allocate-method"]; v != nil {
					allocateIpAddressFromMapToReturn["allocate_method"] = v
				}
				if v := allocateIpAddressFromMap["manual-network"]; v != nil {

					manualNetworkMap := v.(map[string]interface{})

					manualNetworkMapToReturn := make(map[string]interface{})

					if v := manualNetworkMap["name"]; v != nil {
						manualNetworkMapToReturn["name"] = v
					}
					if v := manualNetworkMap["type"]; v != nil {
						manualNetworkMapToReturn["type"] = v
					}
					if v := manualNetworkMap["color"]; v != nil {
						manualNetworkMapToReturn["color"] = v
					}
					if v := manualNetworkMap["domain"]; v != nil {

						domainMap := v.(map[string]interface{})

						domainMapToReturn := make(map[string]interface{})

						if v := domainMap["name"]; v != nil {
							domainMapToReturn["name"] = v
						}
						if v := domainMap["domain-type"]; v != nil {
							domainMapToReturn["domain_type"] = v
						}
						if v := domainMap["uid"]; v != nil {
							domainMapToReturn["uid"] = v
						}

						manualNetworkMapToReturn["domain"] = []interface{}{domainMapToReturn}
					}

					if v := manualNetworkMap["icon"]; v != nil {
						manualNetworkMapToReturn["icon"] = v
					}
					if v := manualNetworkMap["uid"]; v != nil {
						manualNetworkMapToReturn["uid"] = v
					}

					allocateIpAddressFromMapToReturn["manual_network"] = []interface{}{manualNetworkMapToReturn}
				}

				if v := allocateIpAddressFromMap["dhcp-server"]; v != nil {

					dhcpServerMap := v.(map[string]interface{})

					dhcpServerMapToReturn := make(map[string]interface{})

					if v := dhcpServerMap["name"]; v != nil {
						dhcpServerMapToReturn["name"] = v
					}
					if v := dhcpServerMap["type"]; v != nil {
						dhcpServerMapToReturn["type"] = v
					}
					if v := dhcpServerMap["color"]; v != nil {
						dhcpServerMapToReturn["color"] = v
					}
					if v := dhcpServerMap["domain"]; v != nil {

						domainMap := v.(map[string]interface{})

						domainMapToReturn := make(map[string]interface{})

						if v := domainMap["name"]; v != nil {
							domainMapToReturn["name"] = v
						}
						if v := domainMap["domain-type"]; v != nil {
							domainMapToReturn["domain_type"] = v
						}
						if v := domainMap["uid"]; v != nil {
							domainMapToReturn["uid"] = v
						}

						dhcpServerMapToReturn["domain"] = []interface{}{domainMapToReturn}
					}

					if v := dhcpServerMap["icon"]; v != nil {
						dhcpServerMapToReturn["icon"] = v
					}
					if v := dhcpServerMap["uid"]; v != nil {
						dhcpServerMapToReturn["uid"] = v
					}

					allocateIpAddressFromMapToReturn["dhcp_server"] = []interface{}{dhcpServerMapToReturn}
				}

				if v := allocateIpAddressFromMap["virtual-ip-address"]; v != nil {
					allocateIpAddressFromMapToReturn["virtual_ip_address"] = v
				}
				if v := allocateIpAddressFromMap["dhcp-mac-address"]; v != nil {
					allocateIpAddressFromMapToReturn["dhcp_mac_address"] = v
				}
				if v := allocateIpAddressFromMap["optional-parameters"]; v != nil {

					optionalParametersMap := v.(map[string]interface{})

					optionalParametersMapToReturn := make(map[string]interface{})

					if v := optionalParametersMap["use-primary-dns-server"]; v != nil {
						optionalParametersMapToReturn["use_primary_dns_server"] = v
					}
					if v := optionalParametersMap["primary-dns-server"]; v != nil {

						primaryDnsServerMap := v.(map[string]interface{})

						primaryDnsServerMapToReturn := make(map[string]interface{})

						if v := primaryDnsServerMap["name"]; v != nil {
							primaryDnsServerMapToReturn["name"] = v
						}
						if v := primaryDnsServerMap["type"]; v != nil {
							primaryDnsServerMapToReturn["type"] = v
						}
						if v := primaryDnsServerMap["color"]; v != nil {
							primaryDnsServerMapToReturn["color"] = v
						}
						if v := primaryDnsServerMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							primaryDnsServerMapToReturn["domain"] = []interface{}{domainMapToReturn}
						}

						if v := primaryDnsServerMap["icon"]; v != nil {
							primaryDnsServerMapToReturn["icon"] = v
						}
						if v := primaryDnsServerMap["uid"]; v != nil {
							primaryDnsServerMapToReturn["uid"] = v
						}

						optionalParametersMapToReturn["primary_dns_server"] = []interface{}{primaryDnsServerMapToReturn}
					}

					if v := optionalParametersMap["use-first-backup-dns-server"]; v != nil {
						optionalParametersMapToReturn["use_first_backup_dns_server"] = v
					}
					if v := optionalParametersMap["first-backup-dns-server"]; v != nil {

						firstBackupDnsServerMap := v.(map[string]interface{})

						firstBackupDnsServerMapToReturn := make(map[string]interface{})

						if v := firstBackupDnsServerMap["name"]; v != nil {
							firstBackupDnsServerMapToReturn["name"] = v
						}
						if v := firstBackupDnsServerMap["type"]; v != nil {
							firstBackupDnsServerMapToReturn["type"] = v
						}
						if v := firstBackupDnsServerMap["color"]; v != nil {
							firstBackupDnsServerMapToReturn["color"] = v
						}
						if v := firstBackupDnsServerMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							firstBackupDnsServerMapToReturn["domain"] = []interface{}{domainMapToReturn}
						}

						if v := firstBackupDnsServerMap["icon"]; v != nil {
							firstBackupDnsServerMapToReturn["icon"] = v
						}
						if v := firstBackupDnsServerMap["uid"]; v != nil {
							firstBackupDnsServerMapToReturn["uid"] = v
						}

						optionalParametersMapToReturn["first_backup_dns_server"] = []interface{}{firstBackupDnsServerMapToReturn}
					}

					if v := optionalParametersMap["use-second-backup-dns-server"]; v != nil {
						optionalParametersMapToReturn["use_second_backup_dns_server"] = v
					}
					if v := optionalParametersMap["second-backup-dns-server"]; v != nil {

						secondBackupDnsServerMap := v.(map[string]interface{})

						secondBackupDnsServerMapToReturn := make(map[string]interface{})

						if v := secondBackupDnsServerMap["name"]; v != nil {
							secondBackupDnsServerMapToReturn["name"] = v
						}
						if v := secondBackupDnsServerMap["type"]; v != nil {
							secondBackupDnsServerMapToReturn["type"] = v
						}
						if v := secondBackupDnsServerMap["color"]; v != nil {
							secondBackupDnsServerMapToReturn["color"] = v
						}
						if v := secondBackupDnsServerMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							secondBackupDnsServerMapToReturn["domain"] = []interface{}{domainMapToReturn}
						}

						if v := secondBackupDnsServerMap["icon"]; v != nil {
							secondBackupDnsServerMapToReturn["icon"] = v
						}
						if v := secondBackupDnsServerMap["uid"]; v != nil {
							secondBackupDnsServerMapToReturn["uid"] = v
						}

						optionalParametersMapToReturn["second_backup_dns_server"] = []interface{}{secondBackupDnsServerMapToReturn}
					}

					if v := optionalParametersMap["dns-suffixes"]; v != nil {
						optionalParametersMapToReturn["dns_suffixes"] = v
					}
					if v := optionalParametersMap["use-primary-wins-server"]; v != nil {
						optionalParametersMapToReturn["use_primary_wins_server"] = v
					}
					if v := optionalParametersMap["primary-wins-server"]; v != nil {

						primaryWinsServerMap := v.(map[string]interface{})

						primaryWinsServerMapToReturn := make(map[string]interface{})

						if v := primaryWinsServerMap["name"]; v != nil {
							primaryWinsServerMapToReturn["name"] = v
						}
						if v := primaryWinsServerMap["type"]; v != nil {
							primaryWinsServerMapToReturn["type"] = v
						}
						if v := primaryWinsServerMap["color"]; v != nil {
							primaryWinsServerMapToReturn["color"] = v
						}
						if v := primaryWinsServerMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							primaryWinsServerMapToReturn["domain"] = []interface{}{domainMapToReturn}
						}

						if v := primaryWinsServerMap["icon"]; v != nil {
							primaryWinsServerMapToReturn["icon"] = v
						}
						if v := primaryWinsServerMap["uid"]; v != nil {
							primaryWinsServerMapToReturn["uid"] = v
						}

						optionalParametersMapToReturn["primary_wins_server"] = []interface{}{primaryWinsServerMapToReturn}
					}

					if v := optionalParametersMap["use-first-backup-wins-server"]; v != nil {
						optionalParametersMapToReturn["use_first_backup_wins_server"] = v
					}
					if v := optionalParametersMap["first-backup-wins-server"]; v != nil {

						firstBackupWinsServerMap := v.(map[string]interface{})

						firstBackupWinsServerMapToReturn := make(map[string]interface{})

						if v := firstBackupWinsServerMap["name"]; v != nil {
							firstBackupWinsServerMapToReturn["name"] = v
						}
						if v := firstBackupWinsServerMap["type"]; v != nil {
							firstBackupWinsServerMapToReturn["type"] = v
						}
						if v := firstBackupWinsServerMap["color"]; v != nil {
							firstBackupWinsServerMapToReturn["color"] = v
						}
						if v := firstBackupWinsServerMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							firstBackupWinsServerMapToReturn["domain"] = []interface{}{domainMapToReturn}
						}

						if v := firstBackupWinsServerMap["icon"]; v != nil {
							firstBackupWinsServerMapToReturn["icon"] = v
						}
						if v := firstBackupWinsServerMap["uid"]; v != nil {
							firstBackupWinsServerMapToReturn["uid"] = v
						}

						optionalParametersMapToReturn["first_backup_wins_server"] = []interface{}{firstBackupWinsServerMapToReturn}
					}

					if v := optionalParametersMap["use-second-backup-wins-server"]; v != nil {
						optionalParametersMapToReturn["use_second_backup_wins_server"] = v
					}
					if v := optionalParametersMap["second-backup-wins-server"]; v != nil {

						secondBackupWinsServerMap := v.(map[string]interface{})

						secondBackupWinsServerMapToReturn := make(map[string]interface{})

						if v := secondBackupWinsServerMap["name"]; v != nil {
							secondBackupWinsServerMapToReturn["name"] = v
						}
						if v := secondBackupWinsServerMap["type"]; v != nil {
							secondBackupWinsServerMapToReturn["type"] = v
						}
						if v := secondBackupWinsServerMap["color"]; v != nil {
							secondBackupWinsServerMapToReturn["color"] = v
						}
						if v := secondBackupWinsServerMap["domain"]; v != nil {

							domainMap := v.(map[string]interface{})

							domainMapToReturn := make(map[string]interface{})

							if v := domainMap["name"]; v != nil {
								domainMapToReturn["name"] = v
							}
							if v := domainMap["domain-type"]; v != nil {
								domainMapToReturn["domain_type"] = v
							}
							if v := domainMap["uid"]; v != nil {
								domainMapToReturn["uid"] = v
							}

							secondBackupWinsServerMapToReturn["domain"] = []interface{}{domainMapToReturn}
						}

						if v := secondBackupWinsServerMap["icon"]; v != nil {
							secondBackupWinsServerMapToReturn["icon"] = v
						}
						if v := secondBackupWinsServerMap["uid"]; v != nil {
							secondBackupWinsServerMapToReturn["uid"] = v
						}

						optionalParametersMapToReturn["second_backup_wins_server"] = []interface{}{secondBackupWinsServerMapToReturn}
					}

					if v := optionalParametersMap["ip-lease-duration"]; v != nil {
						optionalParametersMapToReturn["ip_lease_duration"] = v
					}

					allocateIpAddressFromMapToReturn["optional_parameters"] = []interface{}{optionalParametersMapToReturn}
				}

				officeModeMapToReturn["allocate_ip_address_from"] = []interface{}{allocateIpAddressFromMapToReturn}
			}

			if v := officeModeMap["support-multiple-interfaces"]; v != nil {
				officeModeMapToReturn["support_multiple_interfaces"] = v
			}
			if v := officeModeMap["perform-anti-spoofing"]; v != nil {
				officeModeMapToReturn["perform_anti_spoofing"] = v
			}
			if v := officeModeMap["anti-spoofing-additional-addresses"]; v != nil {

				antiSpoofingAdditionalAddressesMap := v.(map[string]interface{})

				antiSpoofingAdditionalAddressesMapToReturn := make(map[string]interface{})

				if v := antiSpoofingAdditionalAddressesMap["name"]; v != nil {
					antiSpoofingAdditionalAddressesMapToReturn["name"] = v
				}
				if v := antiSpoofingAdditionalAddressesMap["type"]; v != nil {
					antiSpoofingAdditionalAddressesMapToReturn["type"] = v
				}
				if v := antiSpoofingAdditionalAddressesMap["color"]; v != nil {
					antiSpoofingAdditionalAddressesMapToReturn["color"] = v
				}
				if v := antiSpoofingAdditionalAddressesMap["domain"]; v != nil {

					domainMap := v.(map[string]interface{})

					domainMapToReturn := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToReturn["name"] = v
					}
					if v := domainMap["domain-type"]; v != nil {
						domainMapToReturn["domain_type"] = v
					}
					if v := domainMap["uid"]; v != nil {
						domainMapToReturn["uid"] = v
					}

					antiSpoofingAdditionalAddressesMapToReturn["domain"] = []interface{}{domainMapToReturn}
				}

				if v := antiSpoofingAdditionalAddressesMap["icon"]; v != nil {
					antiSpoofingAdditionalAddressesMapToReturn["icon"] = v
				}
				if v := antiSpoofingAdditionalAddressesMap["uid"]; v != nil {
					antiSpoofingAdditionalAddressesMapToReturn["uid"] = v
				}

				officeModeMapToReturn["anti_spoofing_additional_addresses"] = []interface{}{antiSpoofingAdditionalAddressesMapToReturn}
			}

			vpnSettingsMapToReturn["office_mode"] = []interface{}{officeModeMapToReturn}
		}

		if v := vpnSettingsMap["remote-access"]; v != nil {

			remoteAccessMap := v.(map[string]interface{})

			remoteAccessMapToReturn := make(map[string]interface{})

			if v := remoteAccessMap["support-l2tp"]; v != nil {
				remoteAccessMapToReturn["support_l2tp"] = v
			}
			if v := remoteAccessMap["l2tp-auth-method"]; v != nil {
				remoteAccessMapToReturn["l2tp_auth_method"] = v
			}
			if v := remoteAccessMap["l2tp-certificate"]; v != nil {
				remoteAccessMapToReturn["l2tp_certificate"] = v
			}
			if v := remoteAccessMap["allow-vpn-clients-to-route-traffic"]; v != nil {
				remoteAccessMapToReturn["allow_vpn_clients_to_route_traffic"] = v
			}
			if v := remoteAccessMap["support-nat-traversal-mechanism"]; v != nil {
				remoteAccessMapToReturn["support_nat_traversal_mechanism"] = v
			}
			if v := remoteAccessMap["nat-traversal-service"]; v != nil {

				natTraversalServiceMap := v.(map[string]interface{})

				natTraversalServiceMapToReturn := make(map[string]interface{})

				if v := natTraversalServiceMap["name"]; v != nil {
					natTraversalServiceMapToReturn["name"] = v
				}
				if v := natTraversalServiceMap["type"]; v != nil {
					natTraversalServiceMapToReturn["type"] = v
				}
				if v := natTraversalServiceMap["color"]; v != nil {
					natTraversalServiceMapToReturn["color"] = v
				}
				if v := natTraversalServiceMap["domain"]; v != nil {

					domainMap := v.(map[string]interface{})

					domainMapToReturn := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToReturn["name"] = v
					}
					if v := domainMap["domain-type"]; v != nil {
						domainMapToReturn["domain_type"] = v
					}
					if v := domainMap["uid"]; v != nil {
						domainMapToReturn["uid"] = v
					}

					natTraversalServiceMapToReturn["domain"] = []interface{}{domainMapToReturn}
				}

				if v := natTraversalServiceMap["icon"]; v != nil {
					natTraversalServiceMapToReturn["icon"] = v
				}
				if v := natTraversalServiceMap["uid"]; v != nil {
					natTraversalServiceMapToReturn["uid"] = v
				}

				remoteAccessMapToReturn["nat_traversal_service"] = []interface{}{natTraversalServiceMapToReturn}
			}

			if v := remoteAccessMap["support-visitor-mode"]; v != nil {
				remoteAccessMapToReturn["support_visitor_mode"] = v
			}
			if v := remoteAccessMap["visitor-mode-service"]; v != nil {

				visitorModeServiceMap := v.(map[string]interface{})

				visitorModeServiceMapToReturn := make(map[string]interface{})

				if v := visitorModeServiceMap["name"]; v != nil {
					visitorModeServiceMapToReturn["name"] = v
				}
				if v := visitorModeServiceMap["type"]; v != nil {
					visitorModeServiceMapToReturn["type"] = v
				}
				if v := visitorModeServiceMap["color"]; v != nil {
					visitorModeServiceMapToReturn["color"] = v
				}
				if v := visitorModeServiceMap["domain"]; v != nil {

					domainMap := v.(map[string]interface{})

					domainMapToReturn := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToReturn["name"] = v
					}
					if v := domainMap["domain-type"]; v != nil {
						domainMapToReturn["domain_type"] = v
					}
					if v := domainMap["uid"]; v != nil {
						domainMapToReturn["uid"] = v
					}

					visitorModeServiceMapToReturn["domain"] = []interface{}{domainMapToReturn}
				}

				if v := visitorModeServiceMap["icon"]; v != nil {
					visitorModeServiceMapToReturn["icon"] = v
				}
				if v := visitorModeServiceMap["uid"]; v != nil {
					visitorModeServiceMapToReturn["uid"] = v
				}

				remoteAccessMapToReturn["visitor_mode_service"] = []interface{}{visitorModeServiceMapToReturn}
			}

			if v := remoteAccessMap["visitor-mode-interface"]; v != nil {
				remoteAccessMapToReturn["visitor_mode_interface"] = v
			}

			vpnSettingsMapToReturn["remote_access"] = []interface{}{remoteAccessMapToReturn}
		}

		if v := vpnSettingsMap["saml-portal-settings"]; v != nil {

			samlPortalSettingsMap := v.(map[string]interface{})

			samlPortalSettingsMapToReturn := make(map[string]interface{})

			if v := samlPortalSettingsMap["enabled"]; v != nil {
				samlPortalSettingsMapToReturn["enabled"] = v
			}
			if v := samlPortalSettingsMap["portal-web-settings"]; v != nil {

				portalWebSettingsMap := v.(map[string]interface{})

				portalWebSettingsMapToReturn := make(map[string]interface{})

				if v := portalWebSettingsMap["aliases"]; v != nil {
					portalWebSettingsMapToReturn["aliases"] = v
				}
				if v := portalWebSettingsMap["ip-address"]; v != nil {
					portalWebSettingsMapToReturn["ip_address"] = v
				}
				if v := portalWebSettingsMap["main-url"]; v != nil {
					portalWebSettingsMapToReturn["main_url"] = v
				}

				samlPortalSettingsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
			}

			if v := samlPortalSettingsMap["certificate-settings"]; v != nil {

				certificateSettingsMap := v.(map[string]interface{})

				certificateSettingsMapToReturn := make(map[string]interface{})

				if v := certificateSettingsMap["certificate"]; v != nil {
					certificateSettingsMapToReturn["certificate"] = v
				}
				if v := certificateSettingsMap["certificate-dn"]; v != nil {
					certificateSettingsMapToReturn["certificate_dn"] = v
				}
				if v := certificateSettingsMap["certificate-valid-from"]; v != nil {
					certificateSettingsMapToReturn["certificate_valid_from"] = v
				}
				if v := certificateSettingsMap["certificate-valid-to"]; v != nil {
					certificateSettingsMapToReturn["certificate_valid_to"] = v
				}

				samlPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
			}

			if v := samlPortalSettingsMap["accessibility"]; v != nil {

				accessibilityMap := v.(map[string]interface{})

				accessibilityMapToReturn := make(map[string]interface{})

				if v := accessibilityMap["allow-access-from"]; v != nil {
					accessibilityMapToReturn["allow_access_from"] = v
				}
				if v := accessibilityMap["internal-access-settings"]; v != nil {

					internalAccessSettingsMap := v.(map[string]interface{})

					internalAccessSettingsMapToReturn := make(map[string]interface{})

					if v := internalAccessSettingsMap["undefined"]; v != nil {
						internalAccessSettingsMapToReturn["undefined"] = v
					}
					if v := internalAccessSettingsMap["dmz"]; v != nil {
						internalAccessSettingsMapToReturn["dmz"] = v
					}
					if v := internalAccessSettingsMap["vpn"]; v != nil {
						internalAccessSettingsMapToReturn["vpn"] = v
					}

					accessibilityMapToReturn["internal_access_settings"] = []interface{}{internalAccessSettingsMapToReturn}
				}

				samlPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
			}

			vpnSettingsMapToReturn["saml_portal_settings"] = []interface{}{samlPortalSettingsMapToReturn}
		}

		if v := vpnSettingsMap["vpn-clients"]; v != nil {

			vpnClientsMap := v.(map[string]interface{})

			vpnClientsMapToReturn := make(map[string]interface{})

			if v := vpnClientsMap["enable-endpoint-security-vpn"]; v != nil {
				vpnClientsMapToReturn["enable_endpoint_security_vpn"] = v
			}
			if v := vpnClientsMap["enable-cp-mobile-for-windows"]; v != nil {
				vpnClientsMapToReturn["enable_cp_mobile_for_windows"] = v
			}
			if v := vpnClientsMap["enable-secu-remote"]; v != nil {
				vpnClientsMapToReturn["enable_secu_remote"] = v
			}
			if v := vpnClientsMap["enable-capsule-vpn-connect"]; v != nil {
				vpnClientsMapToReturn["enable_capsule_vpn_connect"] = v
			}
			if v := vpnClientsMap["enable-ssl-network-extender"]; v != nil {
				vpnClientsMapToReturn["enable_ssl_network_extender"] = v
			}
			if v := vpnClientsMap["gateway-authentication-certificate"]; v != nil {
				vpnClientsMapToReturn["gateway_authentication_certificate"] = v
			}

			vpnSettingsMapToReturn["vpn_clients"] = []interface{}{vpnClientsMapToReturn}
		}

		if v := vpnSettingsMap["vpn-domain"]; v != nil {

			vpnDomainMap := v.(map[string]interface{})

			vpnDomainMapToReturn := make(map[string]interface{})

			if v := vpnDomainMap["name"]; v != nil {
				vpnDomainMapToReturn["name"] = v
			}
			if v := vpnDomainMap["type"]; v != nil {
				vpnDomainMapToReturn["type"] = v
			}
			if v := vpnDomainMap["color"]; v != nil {
				vpnDomainMapToReturn["color"] = v
			}
			if v := vpnDomainMap["domain"]; v != nil {

				domainMap := v.(map[string]interface{})

				domainMapToReturn := make(map[string]interface{})

				if v := domainMap["name"]; v != nil {
					domainMapToReturn["name"] = v
				}
				if v := domainMap["domain-type"]; v != nil {
					domainMapToReturn["domain_type"] = v
				}
				if v := domainMap["uid"]; v != nil {
					domainMapToReturn["uid"] = v
				}

				vpnDomainMapToReturn["domain"] = []interface{}{domainMapToReturn}
			}

			if v := vpnDomainMap["icon"]; v != nil {
				vpnDomainMapToReturn["icon"] = v
			}
			if v := vpnDomainMap["uid"]; v != nil {
				vpnDomainMapToReturn["uid"] = v
			}

			vpnSettingsMapToReturn["vpn_domain"] = []interface{}{vpnDomainMapToReturn}
		}

		if v := vpnSettingsMap["vpn-domain-exclude-external-ip-addresses"]; v != nil {
			vpnSettingsMapToReturn["vpn_domain_exclude_external_ip_addresses"] = v
		}
		if v := vpnSettingsMap["vpn-domain-type"]; v != nil {
			vpnSettingsMapToReturn["vpn_domain_type"] = v
		}
		if v := vpnSettingsMap["enable-clientless-vpn"]; v != nil {
			vpnSettingsMapToReturn["enable_clientless_vpn"] = v
		}
		if v := vpnSettingsMap["clientless-vpn-settings"]; v != nil {

			clientlessVpnSettingsMap := v.(map[string]interface{})

			clientlessVpnSettingsMapToReturn := make(map[string]interface{})

			if v := clientlessVpnSettingsMap["certificate-gateway-authentication"]; v != nil {
				clientlessVpnSettingsMapToReturn["certificate_gateway_authentication"] = v
			}
			if v := clientlessVpnSettingsMap["client-authentication"]; v != nil {
				clientlessVpnSettingsMapToReturn["client_authentication"] = v
			}
			if v := clientlessVpnSettingsMap["concurrent-servers-or-processes"]; v != nil {
				clientlessVpnSettingsMapToReturn["concurrent_servers_or_processes"] = v
			}
			if v := clientlessVpnSettingsMap["accept-only-3des"]; v != nil {
				clientlessVpnSettingsMapToReturn["accept_only_3des"] = v
			}

			vpnSettingsMapToReturn["clientless_vpn_settings"] = []interface{}{clientlessVpnSettingsMapToReturn}
		}

		_ = d.Set("vpn_settings", []interface{}{vpnSettingsMapToReturn})

	} else {
		_ = d.Set("vpn_settings", nil)
	}

	if v := cluster["tags"]; v != nil {
		tagsJson := v.([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			for _, tag := range tagsJson {
				tagsIds = append(tagsIds, tag.(map[string]interface{})["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if v := cluster["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := cluster["color"]; v != nil {
		_ = d.Set("color", v)
	}

	return nil
}
