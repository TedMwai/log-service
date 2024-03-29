import { Metadata } from "next";
import { Activity, Bug, Timer } from "lucide-react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

import { columns } from "./components/columns";
import { DataTable } from "./components/data-table";
import { LogSchema } from "./data/schema";


export const metadata: Metadata = {
  title: "Dashboard",
  description: "Dashboard & log tracker built using Tanstack Table.",
}

async function getLogs(): Promise<LogSchema[]> {
  const response = await fetch("http://localhost:8080/logs")
  return response.json()
}

export default async function HomePage() {
  const logs = await getLogs()

  return (
    <>
      <div className="md:hidden"></div>
      <div className="hidden flex-col md:flex">
        <div className="flex-1 space-y-4 p-8 pt-6">
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
            <Card>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">
                  Number of Logs
                </CardTitle>
                <Activity className="h-4 w-4 text-muted-foreground" />
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold pt-2.5">3,000</div>
              </CardContent>
            </Card>
            <Card>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">
                  Number of Fatal Errors
                </CardTitle>
                <Bug className="h-4 w-4 text-muted-foreground" />
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold pt-2.5">30</div>
              </CardContent>
            </Card>
            <Card>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">
                  Average Response Time
                </CardTitle>
                <Timer className="h-4 w-4 text-muted-foreground" />
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold pt-2.5">2.4 ms</div>
              </CardContent>
            </Card>
          </div>
          <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
            <div className="flex items-center justify-between space-y-2">
              <div>
                <h2 className="text-2xl font-bold tracking-tight">
                  Here are your logs!
                </h2>
              </div>
            </div>
            <DataTable data={logs} columns={columns} />
          </div>
        </div>
      </div>
    </>
  )
}