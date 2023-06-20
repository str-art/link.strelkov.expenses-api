import ExpenseForm from "@/components/ExpenseForm";

export default async function Home() {
  const response = await fetch(
    "https://w2llq9lslg.execute-api.eu-central-1.amazonaws.com/dev/categories",
    { cache: "no-cache" }
  );
  const categories = (await response.json()) as Array<Category>;

  return (
    <>
      <ExpenseForm categories={categories.map(({ name }) => name)} />
    </>
  );
}
