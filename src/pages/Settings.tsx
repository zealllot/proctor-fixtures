import React from 'react';

export default function Settings() {
  const save = (e: React.FormEvent) => { e.preventDefault(); };
  return (
    <form onSubmit={save}>
      <input name="email" />
      <button>Save</button>
    </form>
  );
}
