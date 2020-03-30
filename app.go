func (a *App) initRepo() error {
	log.Print("Initializing repository")

	err := os.MkdirAll(a.repoDir, 0755)
	// Check err

	cmd := exec.Command("git", "--git-dir="+a.repoDir, "init")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	// Check err

	cmd = exec.Command("git", "--git-dir="+a.repoDir, "remote", "add", "origin", fmt.Sprintf("git@github.com:%s.git", a.Repo))
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	// Check err

	return nil
}
