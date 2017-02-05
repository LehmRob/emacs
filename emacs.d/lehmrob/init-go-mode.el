(maybe-require-package 'go-mode)

(defun my-go-mode-hook ()
  (setq gofmt-command (concat (getenv "GOPATH") ,"/bin/goimports"))
  (add-hook 'before-save-hook 'gofmt-before-save)
  (if (not (string-match "go" compile-command))
      (set (make-local-variable 'compile-command)
           "go build -v && go test -v && go vet"))
  (local-set-key (kbd "M-.") 'godef-jump))
(add-hook 'go-mode-hook 'my-go-mode-hook)
(provide 'init-go-mode)

