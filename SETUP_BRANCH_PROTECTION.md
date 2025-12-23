# Branch Protection Setup Guide

## 🔒 Enable Automatic Code Quality Enforcement

To make code quality checks **truly mandatory** and **automatic**, configure branch protection rules on GitHub.

---

## Setup Steps (One-Time, Repo Admin Only)

### 1. Go to Repository Settings

Navigate to: `https://github.com/YOUR_USERNAME/go-task-api/settings/branches`

### 2. Add Branch Protection Rule

Click: **"Add rule"** or **"Add branch protection rule"**

### 3. Configure Protection for Main Branch

**Branch name pattern:** `main`

Enable these settings:

#### ✅ Require a pull request before merging
- Require approvals: `1` (optional, adjust as needed)

#### ✅ Require status checks to pass before merging
- **Check:** `Enforce Code Quality Standards` ✓
- **Check:** `code-quality-check` ✓
- ✅ Require branches to be up to date before merging

#### ✅ Do not allow bypassing the above settings
- This ensures even admins must pass checks

### 4. Save Changes

Click: **"Create"** or **"Save changes"**

---

## What This Does

### Before Branch Protection:
```
Developer pushes bad code
  ↓
GitHub Actions runs and fails ❌
  ↓
Developer can still merge to main (BAD!)
```

### After Branch Protection:
```
Developer pushes bad code
  ↓
GitHub Actions runs and fails ❌
  ↓
Merge button is DISABLED (GOOD!)
  ↓
Developer must fix code before merging ✅
```

---

## Result

✅ **Automatic enforcement** - no local setup required  
✅ **Cannot be bypassed** - even by repo admins  
✅ **Blocks bad code** - merge button disabled until checks pass  
✅ **Zero configuration** needed by developers  

---

## Optional: Protect All Branches

To enforce checks on **all branches** (not just main):

**Branch name pattern:** `**` (matches all branches)

Then enable the same settings as above.

---

## Verification

After setup, try pushing code with trailing whitespace:

1. The GitHub Actions workflow will fail ❌
2. The PR will show: "Some checks were not successful"
3. The merge button will be disabled
4. Developer must fix the code to proceed

---

## Questions?

- GitHub Docs: https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches
- Atlassian Git Hooks: https://www.atlassian.com/git/tutorials/git-hooks

