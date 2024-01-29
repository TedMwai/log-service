import themes from "@/themes/index.json"
import { ChevronLeft, ChevronRight } from "lucide-react"

import { sanitizeName } from "@/lib/utils"

export default function PageHeader({ theme }: { theme: string }) {
  const currentTheme =
    themes.find((t) => sanitizeName(t.name) === sanitizeName(theme)) ??
    themes[0]

  return (
    <section className="flex min-h-[120px] w-full flex-col items-center justify-between pb-2 pt-6 sm:flex-row md:min-h-[100px] md:py-5 md:pb-8">
      <div className="flex max-w-[980px] flex-col items-start gap-2">
        <h1 className="text-2xl font-extrabold leading-tight tracking-tighter md:text-4xl">
          Logs
          <span
            className={
              "pl-2 text-base font-light leading-normal tracking-wide md:text-xl"
            }
          >
            by group IV
          </span>
        </h1>
      </div>
    </section>
  )
}
