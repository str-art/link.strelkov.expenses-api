import { Axios } from "axios";

export default async function addExpense(client: Axios, expense: ExpenseInput) {
  console.log(expense);
  await client.post(`/expenses/${expense.category}/${expense.date}`, {
    amount: expense.amount,
    name: expense.name,
  });
}
