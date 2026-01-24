import { LoginSchema } from "@/types/auth";
import z from "zod";

export const LoginForm : z.ZodType<LoginSchema> = z.object({
  email: z
    .email()
    .min(1, { message: 'Email là bắt buộc' }),
  password: z
    .string()
    .min(1, { message: 'Mật khẩu là bắt buộc' })
    .min(6, { message: 'Mật khẩu phải có ít nhất 6 ký tự' }),
});