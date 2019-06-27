package account_check_printing

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AccountJournal().DeclareModel()

	h.AccountJournal().Methods().ComputeCheckPrintingPaymentMethodSelected().DeclareMethod(
		`ComputeCheckPrintingPaymentMethodSelected`,
		func(rs m.AccountJournalSet) {
			//        self.check_printing_payment_method_selected = any(
			//            pm.code == 'check_printing' for pm in self.outbound_payment_method_ids)
		})
	h.AccountJournal().Methods().GetCheckNextNumber().DeclareMethod(
		`GetCheckNextNumber`,
		func(rs m.AccountJournalSet) {
			//        if self.check_sequence_id:
			//            self.check_next_number = self.check_sequence_id.number_next_actual
			//        else:
			//            self.check_next_number = 1
		})
	h.AccountJournal().Methods().SetCheckNextNumber().DeclareMethod(
		`SetCheckNextNumber`,
		func(rs m.AccountJournalSet) {
			//        if self.check_next_number < self.check_sequence_id.number_next_actual:
			//            raise ValidationError(_("The last check number was %s. In order to avoid a check being rejected "
			//                                    "by the bank, you can only use a greater number.") % self.check_sequence_id.number_next_actual)
			//        if self.check_sequence_id:
			//            self.check_sequence_id.sudo().number_next_actual = self.check_next_number
		})
	h.AccountJournal().AddFields(map[string]models.FieldDefinition{
		"CheckManualSequencing": models.BooleanField{
			String:  "Manual Numbering",
			Default: models.DefaultValue(false),
			Help:    "Check this option if your pre-printed checks are not numbered.",
		},
		"CheckSequenceId": models.Many2OneField{
			RelationModel: h.IrSequence(),
			String:        "Check Sequence",
			ReadOnly:      true,
			NoCopy:        true,
			Help:          "Checks numbering sequence.",
		},
		"CheckNextNumber": models.IntegerField{
			String:  "Next Check Number",
			Compute: h.AccountJournal().Methods().GetCheckNextNumber(),
			//inverse='_set_check_next_number'
			Help: "Sequence number of the next printed check.",
		},
		"CheckPrintingPaymentMethodSelected": models.BooleanField{
			Compute: h.AccountJournal().Methods().ComputeCheckPrintingPaymentMethodSelected(),
			Help: "Technical feature used to know whether check printing was" +
				"enabled as payment method.",
		},
	})
	h.AccountJournal().Methods().Create().Extend(
		`Create`,
		func(rs m.AccountJournalSet, vals models.RecordData) {
			//        rec = super(AccountJournal, self).create(vals)
			//        if not rec.check_sequence_id:
			//            rec._create_check_sequence()
			//        return rec
		})
	h.AccountJournal().Methods().Copy().Extend(
		`Copy`,
		func(rs m.AccountJournalSet, defaultName models.RecordData) {
			//        rec = super(AccountJournal, self).copy(default)
			//        rec._create_check_sequence()
			//        return rec
		})
	h.AccountJournal().Methods().CreateCheckSequence().DeclareMethod(
		` Create a check sequence for the journal `,
		func(rs m.AccountJournalSet) {
			//        self.check_sequence_id = self.env['ir.sequence'].sudo().create({
			//            'name': self.name + _(" : Check Number Sequence"),
			//            'implementation': 'no_gap',
			//            'padding': 5,
			//            'number_increment': 1,
			//            'company_id': self.company_id.id,
			//        })
		})
	h.AccountJournal().Methods().DefaultOutboundPaymentMethods().DeclareMethod(
		`DefaultOutboundPaymentMethods`,
		func(rs m.AccountJournalSet) {
			//        methods = super(
			//            AccountJournal, self)._default_outbound_payment_methods()
			//        return methods + self.env.ref('account_check_printing.account_payment_method_check')
		})
	h.AccountJournal().Methods().EnableCheckPrintingOnBankJournals().DeclareMethod(
		` Enables check printing payment method and add a check
sequence on bank journals.
            Called upon module installation via data file.
        `,
		func(rs m.AccountJournalSet) {
			//        check_printing = self.env.ref(
			//            'account_check_printing.account_payment_method_check')
			//        bank_journals = self.search([('type', '=', 'bank')])
			//        for bank_journal in bank_journals:
			//            bank_journal._create_check_sequence()
			//            bank_journal.write({
			//                'outbound_payment_method_ids': [(4, check_printing.id, None)],
			//            })
		})
	h.AccountJournal().Methods().GetJournalDashboardDatas().DeclareMethod(
		`GetJournalDashboardDatas`,
		func(rs m.AccountJournalSet) {
			//        domain_checks_to_print = [
			//            ('journal_id', '=', self.id),
			//            ('payment_method_id.code', '=', 'check_printing'),
			//            ('state', '=', 'posted')
			//        ]
			//        return dict(
			//            super(AccountJournal, self).get_journal_dashboard_datas(),
			//            num_checks_to_print=len(
			//                self.env['account.payment'].search(domain_checks_to_print))
			//        )
		})
	h.AccountJournal().Methods().ActionChecksToPrint().DeclareMethod(
		`ActionChecksToPrint`,
		func(rs m.AccountJournalSet) {
			//        return {
			//            'name': _('Checks to Print'),
			//            'type': 'ir.actions.act_window',
			//            'view_mode': 'list,form,graph',
			//            'res_model': 'account.payment',
			//            'context': dict(
			//                self.env.context,
			//                search_default_checks_to_send=1,
			//                journal_id=self.id,
			//                default_journal_id=self.id,
			//                default_payment_type='outbound',
			//                default_payment_method_id=self.env.ref('account_check_printing.account_payment_method_check').id),
			//        }
		})
}
