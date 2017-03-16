(defun my-go-mode-hook ()
  (setq gofmt-command "/home/rd7/work/go/bin/goimports")
  (add-hook 'before-save-hook 'gofmt-before-save))
(add-hook 'go-mode-hook 'my-go-mode-hook)
(provide 'init-go-mode)

