"use client";

import Stack from "@mui/material/Stack";
import Card from "@mui/material/Card";
import { usePathname } from "next/navigation";

export default function ExpensesList({
  expenses,
}: {
  expenses: Array<Expense>;
}) {
  const pathname = usePathname();

  console.log(pathname);

  return (
    <Stack component={"ol"} spacing={4}>
      {expenses.map((expense) => (
        <Card component={"li"} key={expense.id}>
          {expense.id}
          <br />
          {expense.name}
          <br />
          {expense.date}
        </Card>
      ))}
    </Stack>
  );
}
