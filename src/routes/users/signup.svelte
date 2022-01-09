<form method="post" action="/users.json">
	<h1>Sign Up!</h1>
	<div>
		<label for="username">Username:</label>
		<input name="username" id="username" />
	</div>

	<div>
		<label for="email">Email:</label>
		<input name="email" id="email" />
	</div>

	<div>
		<button type="submit">Sign up!</button>
	</div>
</form>

<style>
	* {
		box-sizing: border-box;
		font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
	}
	h1 {
		margin: 0.5rem auto;
		letter-spacing: 2px;
	}
	form {
		display: flex;
		flex-direction: column;
		width: 18rem;
		margin: auto;
	}
	form > div {
		display: flex;
		justify-content: space-between;
	}
	form > div + * {
		margin-top: 10px;
	}
	form button {
		width: 100%;
	}
</style>
