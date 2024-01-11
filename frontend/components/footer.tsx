import { Copyright } from "lucide-react"

export default function Footer() {
  return (
    <footer className="container inset-x-0 bottom-0">
      <div className="flex justify-between">
        <div className="flex items-end"></div>
        <div className="self-end pb-1">
          <div className="flex items-end text-sm">
            <Copyright height={17} width={17} />
            &nbsp;Log Management by Group 4 {new Date().getFullYear()}.
          </div>
        </div>
        <div className="flex items-end"></div>
      </div>
    </footer>
  )
}
