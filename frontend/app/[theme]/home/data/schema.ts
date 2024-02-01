import { string, z } from "zod"

// We're keeping a simple non-relational schema here.
// IRL, you will have a schema for your data models.
export const taskSchema = z.object({
  ID: z.string(),
  CreatedBy: z.string(),
  UpdatedBy: z.string(),
  CreatedAt: z.string(),
  UpdatedAt: z.string(),
  DeletedAt: z.string(),
  Metadata: z.record(z.any()),
  MicroserviceID: z.string(),
  LogLevel: z.string(),
  Message: z.string(),
})

export type Task = z.infer<typeof taskSchema>
