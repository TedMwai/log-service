import {
  BellRing,
  Bug,
  CircleDollarSign,
  Info,
  ShieldAlert,
  Truck,
  XCircle,
} from "lucide-react"

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
    value: "info",
    label: "INFO",
    icon: Info,
  },
  {
    value: "debug",
    label: "DEBUG",
    icon: Bug,
  },
  {
    value: "error",
    label: "ERROR",
    icon: ShieldAlert,
  },
  {
    value: "fatal",
    label: "FATAL",
    icon: XCircle,
  },
]

export const microservices = [
  {
    label: "Payments",
    value: "payments",
    icon: CircleDollarSign,
  },
  {
    label: "Notifications",
    value: "notifications",
    icon: BellRing,
  },
  {
    label: "Orders",
    value: "orders",
    icon: Truck,
  },
]
