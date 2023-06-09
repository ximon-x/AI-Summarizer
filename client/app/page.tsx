"use client";

import { FormEvent, useState } from "react";

export default function Home() {
  const [summary, setSummary] = useState("");

  const handleFormSubmit = async (event: FormEvent) => {
    event.preventDefault();
    const form = event.target as HTMLFormElement;

    // Get data from the form.
    const formData = {
      content: form.text.value as string,
    };

    const response = await fetch("http://localhost:8000/summary", {
      body: JSON.stringify(formData),
      headers: {
        "Content-Type": "application/json",
      },
      method: "POST",
    });

    const data = await response.json();
    setSummary(data.summary);
  };

  return (
    <div className="flex min-h-screen min-w-fit items-center justify-center">
      <div className="flex h-auto w-auto flex-col rounded-xl bg-white p-4 dark:bg-black">
        <header className="m-4 flex justify-center">
          <h1 className="mx-8 text-4xl font-bold">AI Summarizer</h1>
        </header>
        <main className="flex flex-col items-center justify-between p-12">
          <form
            onSubmit={handleFormSubmit}
            className="flex flex-col items-center justify-between"
          >
            <textarea
              name="text"
              className="w-96 h-96 p-4 border-2 border-gray-300 rounded-xl dark:text-black"
              placeholder="Enter text here..."
              required
            ></textarea>
            <button className="mt-4 w-[100%] h-12 bg-yellow-500 text-white font-bold rounded-xl dark:bg-blue-500">
              Simplify
            </button>
          </form>
          {summary !== "" ? (
            <p className="mt-4 text-center font-bold ">{summary}</p>
          ) : null}
        </main>
        <footer>
          <p className="text-center text-gray-500">
            Powered by <a href="https://platform.openai.com/overview">OpenAI</a>
          </p>
        </footer>
      </div>
    </div>
  );
}
