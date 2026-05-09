import React from 'react';

export default function Settings() {
  const save = (e: React.FormEvent) => { e.preventDefault(); };
  return (
    <form onSubmit={save}>
      <input name="email" />
      <label>Display name
        <input
          name="display_name"
          maxLength={100}
          data-testid="display-name"
        />
      </label>
      <button>Save</button>
    </form>
  );
}
