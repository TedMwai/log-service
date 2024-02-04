import { BellRing, Bug, CircleDollarSign, Info, ShieldAlert, Truck, XCircle } from "lucide-react";





export const labels = [
  {
    value: "bug",
    label: "Bug",
  },
  {
    value: "feature",
    label: "Feature",
  },
  {
    value: "documentation",
    label: "Documentation",
  },
]

export const statuses = [
  {
    value: "INFO",
    label: "INFO",
    icon: Info,
  },
  {
    value: "DEBUG",
    label: "DEBUG",
    icon: Bug,
  },
  {
    value: "ERROR",
    label: "ERROR",
    icon: ShieldAlert,
  },
  {
    value: "FATAL",
    label: "FATAL",
    icon: XCircle,
  },
]

export const microservices = [
  {
    label: "Billing",
    value: "Billing",
    icon: CircleDollarSign,
  },
  {
    label: "Notification",
    value: "Notification",
    icon: BellRing,
  },
  {
    label: "Orders",
    value: "Orders",
    icon: Truck,
  },
]