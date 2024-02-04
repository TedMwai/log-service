import { z } from "zod"

export type LogSchema = {
  CreatedAt: string
  CreatedBy: string
  DeletedAt: string
  ID: string
  LogLevel: string
  Message: string
  Metadata: Record<string, unknown>
  Microservice: {
    CreatedAt: string
    CreatedBy: string
    DeletedAt: string
    Description: string
    ID: string
    Metadata: Record<string, unknown>
    Name: string
    UpdatedAt: string
    UpdatedBy: string
  }
  MicroserviceID: string
  UpdatedAt: string
  UpdatedBy: string
}

