'use client';

import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Field, FieldGroup, FieldLabel } from '@/components/ui/field';
import { Input } from '@/components/ui/input';
import { useState } from 'react';

export function LoginForm({ className, ...props }: React.ComponentProps<'div'>) {
  const [hasError, setHasError] = useState(false);
  const [data, setData] = useState({ field: '', message: '' });

  const handlerSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    console.log('login form');
    setHasError(true);

    let test = {
        email: "super-admin@gmail.com",
        password: "password123"
    }

    try {
      const res = await fetch('http://localhost:8808/api/v1/login', {
        method: 'POST',
        body: JSON.stringify(test),
        headers: { 'Content-Type': 'application/json' },
      });

      if (!res.ok) {
        throw new Error('Đã có lỗi xảy ra từ máy chủ!');
      }

      const data = await res.json();
      console.log('Kết quả:', data);
    } catch (error) {
        console.log('Kết quả lỗi:', data);
    }
  };

  return (
    <div className={cn('flex flex-col gap-6', className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle>Login to your account</CardTitle>
          <CardDescription>Enter your email below to login to your account</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handlerSubmit} noValidate>
            <FieldGroup>
              <Field data-invalid={hasError} className="gap-1">
                <FieldLabel htmlFor="email">Email</FieldLabel>
                <Input
                  id="email"
                  type="email"
                  autoComplete="email"
                  placeholder="m@example.com"
                  aria-invalid={hasError}
                />
              </Field>
              <Field data-invalid={hasError} className="gap-1">
                <FieldLabel htmlFor="password">Password</FieldLabel>
                <Input
                  id="password"
                  type="password"
                  autoComplete="password"
                  aria-invalid={hasError}
                />
                {hasError}
              </Field>
              <Field>
                <Button type="submit">Login</Button>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
