type Expense = {
  id: string;
  amount: number;
  date: string;
  name: string;
};

type Category = {
  name: string;
};

type ExpenseInput = {
  category: string;
  date: string;
  name: string;
  amount: number;
};
