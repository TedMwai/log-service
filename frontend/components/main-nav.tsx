"use client"

import Image from "next/image"
import Link from "next/link"
import logo from "@/images/logo.png"

import { siteConfig } from "@/config/site"
import { sanitizeName } from "@/lib/utils"

import { useTheme } from "./shadcn-theme-provider"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog"
import { useToast } from "./ui/use-toast"

export function copyToClipboard(value: string) {
  navigator.clipboard.writeText(value)
}

export function MainNav() {
  return (
    <div className="flex gap-6 md:gap-10">
      <div className="flex items-center space-x-2">
        <span role="img" aria-label="logo" style={{ fontSize: "45px" }}>
          🪵
        </span>
        <span className="hidden font-bold md:inline-block">
          Applications Log Management
        </span>
      </div>
    </div>
  )
}
